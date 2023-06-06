package openstack

import (
	"goku/component"
	"goku/pkg/netbox"
	"goku/utils/logging"
	"goku/utils/options"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/diskconfig"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/extendedstatus"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/pagination"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	timeFormart      = "2006-01-02 15:04:05"
	mongoTimeFormart = "2006-01-02"
)

type Port struct {
	ClusterName string `json:"cluster_name" bson:"cluster_name"`
	ID          string `json:"id" bson:"id"`
	NetworkID   string `json:"network_id" bson:"network_id"`
	MacAddress  string `json:"mac_address" bson:"mac_address"`
	IP          string `json:"ip" bson:"ip"`
	DeviceOwner string `json:"device_owner" bson:"device_owner"`
	DeviceID    string `json:"device_id" bson:"device_id"`
}

type sync interface {
	SyncData()
	Auth()
}

type OpenStack struct {
	component.AuthInfo
}

type ServerWithExt struct {
	servers.Server
	availabilityzones.ServerAvailabilityZoneExt
	extendedstatus.ServerExtendedStatusExt
	diskconfig.ServerDiskConfigExt
}

type ProjectMap struct {
	ClusterName string            `json:"cluster_name" bson:"cluster_name"`
	ProjectsMap map[string]string `json:"projects_map" bson:"projects_map"`
}

type Flavor struct {
	Name  string `json:"name"`
	RAM   int    `json:"ram"`
	VCPUs int    `json:"vcpus"`
	Disk  int    `json:"disk"`
}

type Server struct {
	Name        string   `bson:"name" json:"name"`
	ClusterName string   `bson:"cluster_name" json:"cluster_name"`
	Project     string   `bson:"project" json:"project"`
	Uuid        string   `bson:"uuid" json:"uuid"`
	Status      string   `bson:"status" json:"status"`
	IP          []string `bson:"ip" json:"ip"`
	Flavor      `bson:"flavor" json:"flavor"`
	VolumesSize int    `json:"volumes_size" bson:"volumes_size"`
	Node        string `bson:"host" json:"node"`
	Created     string `bson:"created" json:"created"`
	New         bool   `bson:"new" json:"new"`
	NetboxVMID  int64  `bson:"netbox_vmid" json:"netbox_vmid"`
}

// 开始同步数据
func Do(ClusterList []component.AuthInfo) {
	//初始化projectMap
	//创建链接
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()
	_, err = cli.DB("openstack").C("project").RemoveAll(nil)
	_, err = cli.DB("openstack").C("port").RemoveAll(nil)
	if err != nil {
		logging.Error("清空集合失败 Err:", err)
	}
	for _, cluster := range ClusterList {
		clusterInfo := OpenStack{cluster}
		//clusterInfo.SyncDataToMongo(mongo)
		clusterInfo.SyncDataToMongo(cli)
	}
}

// 从netbox删除退库的机器
func DeleteVMFromNetbox(ClusterList []component.AuthInfo) {
	//创建mongodb 链接
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()

	//获取vmMap
	VMMap, err := netbox.GetVMMap()
	if err != nil {
		logging.Error("netbox 连接发生错误", err)
		return
	}
	logging.Info(len(VMMap))
	//新旧对比
	var news []Server
	var oldSs []Server

	CnameTime := time.Now().Format(mongoTimeFormart)
	h, _ := time.ParseDuration("-24h")
	t := time.Now().Add(h)
	yesterdayCnameTime := t.Format(mongoTimeFormart)
	for _, cluster := range ClusterList {
		DelVM := &netbox.DelVM{}
		oldSs = nil
		client := cli.DB("openstack").C(cluster.ClusterName + yesterdayCnameTime)
		_ = client.Find(nil).All(&oldSs)
		for _, s := range oldSs {
			news = nil
			client := cli.DB("openstack").C(cluster.ClusterName + CnameTime)
			_ = client.Find(bson.M{"uuid": s.Uuid}).All(&news)
			if len(news) == 0 {
				DelVM.ID = VMMap[s.Uuid]
				DelVM.IPS = s.IP
				// netbox.DueVM <- DelVM
				netbox.DeleteVMFromNetbox(DelVM)
			}
		}
	}
}

// 写数据到mongo
func (o *OpenStack) SyncDataToMongo(cli *mgo.Session) {
	method := "SyncDataToMongo"
	logging.Info("begin", method)
	//var actual []ServerWithExt
	//获取认证
	provider, err := o.Auth()
	if err != nil {
		logging.Error(o.ClusterName, "认证失败:", err)
		return
	}
	//获取nodeList 并且创建ID->Name map
	nodeList, err := NodeList(provider)
	if err != nil {
		logging.Error(err)
	}
	nodeMap := map[string]string{}
	for _, node := range nodeList {
		nodeMap[node.HypervisorHostname] = node.HostIP
	}
	//获取flavorList 并且创建ID->Name map
	flavorList := ListFlavor(provider)
	flavorMap := map[string]Flavor{}
	for _, flaovr := range flavorList {
		flavorMap[flaovr.ID] = Flavor{flaovr.Name, flaovr.RAM, flaovr.VCPUs, flaovr.Disk}
	}

	//获取projectList 并且创建ID->Name map
	projectList := ListProject(provider)
	projectsMap := map[string]string{}
	for _, project := range projectList {
		projectsMap[project.ID] = project.Name
	}
	projectClient := cli.DB("openstack").C("project")
	cErr := projectClient.Insert(ProjectMap{o.ClusterName, projectsMap})
	if cErr != nil {
		logging.Error("insert mongo error:", cErr)
	}

	//获取AllPorts
	port := Port{}
	portList := ListPort(provider)
	portClient := cli.DB("openstack").C("port")
	for _, p := range portList {
		port = Port{
			ClusterName: o.ClusterName,
			ID:          p.ID,
			NetworkID:   p.NetworkID,
			MacAddress:  p.MACAddress,
			DeviceOwner: p.DeviceOwner,
			DeviceID:    p.DeviceID,
		}
		if len(p.FixedIPs) != 0 {
			port.IP = p.FixedIPs[0].IPAddress
		}
		cErr := portClient.Insert(port)
		if cErr != nil {
			logging.Error("insert mongo error:", cErr)
		}
	}

	//获取serverList
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	opts := servers.ListOpts{AllTenants: true}
	allPages, err := servers.List(client, opts).AllPages()
	//解析返回值
	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		logging.Error("extract response data error:", err)
		return
	}
	//生成集合日期
	Cname := time.Now().Format(mongoTimeFormart)

	//清空数据源
	_, err = cli.DB("openstack").C(o.ClusterName + Cname).RemoveAll(nil)
	if err != nil {
		logging.Error("清空集合%s失败 Err:", o.ClusterName, err)
	}

	//生成数据和同步数据
	serverClient := cli.DB("openstack").C(o.ClusterName + Cname)
	for _, item := range allServers {
		if err != nil {
			logging.Debug("获取 console 失败：", err)
		}
		var Addrs []string
		for _, addrMap := range item.Addresses {
			addr := addrMap[0]["addr"]
			value, ok := addr.(string)
			if !ok {
				logging.Error("It's not ok for type string")
				return
			}
			Addrs = append(Addrs, value)
		}
		var VolumesSize int
		if len(item.AttachedVolumes) != 0 {
			VolumesSize = GetVolumesSizeV2(provider, item.AttachedVolumes)
		} else {
			VolumesSize = flavorMap[item.Flavor["id"].(string)].Disk
		}
		var New bool
		h, _ := time.ParseDuration("-24h")
		t := time.Now().Add(h)
		if item.Created.Sub(t) >= 1 {
			New = true
		} else {
			New = false
		}
		data := Server{
			Name:        item.Name,
			ClusterName: o.ClusterName,
			Project:     projectsMap[item.TenantID],
			Uuid:        item.ID,
			Status:      item.Status,
			IP:          Addrs,
			Flavor:      flavorMap[item.Flavor["id"].(string)],
			VolumesSize: VolumesSize,
			Node:        item.HostName,
			Created:     time.Time(item.Created).Format(timeFormart),
			New:         New,
		}
		if New && options.Conf.Netbox.Onoff {
			SyncVMToNETBOX(data)
		}
		cErr := serverClient.Insert(data)
		if cErr != nil {
			logging.Error("insert mongo error:", cErr)
		}
	}
}

// 写数据到netbox
func SyncVMToNETBOX(S Server) {
	logging.Info("同步netbox:", S.Name)
	tenant := strings.Split(S.Project, "-")
	logging.Debug(tenant)
	vm := &netbox.Vm{
		Name:        S.Name,
		ClusterName: S.ClusterName,
		Vcpus:       float64(S.VCPUs),
		Memory:      int64(S.RAM),
		Disk:        int64(S.Disk),
		Uuid:        S.Uuid,
		Vmcreated:   S.Created,
		Host:        S.Node,
		Tenant:      tenant[0],
	}
	if len(S.IP) > 0 {
		vm.Eth0 = S.IP[0]
	}
	if len(S.IP) > 1 {
		vm.Eth1 = S.IP[1]
	}
	if len(S.IP) > 2 {
		vm.Eth2 = S.IP[2]
	}
	if len(S.IP) > 3 {
		vm.Eth3 = S.IP[3]
	}
	//if NetboxVMID, _ := netbox.SyncVM2Netbox(vm); NetboxVMID != 500000 {
	//	return NetboxVMID
	//}else {
	//	return 500000
	//}
	//netbox.NewVM <- vm
	netbox.SyncVM2Netbox(vm)
}

// 获取宿主机名称
func NodeList(provider *gophercloud.ProviderClient) (nodeList []hypervisors.Hypervisor, err error) {
	method := "NodeList"
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, "认证失败: %v", err)
	}
	allPages, err := hypervisors.List(client, nil).AllPages()
	if err != nil {
		panic(err)
	}
	allNodes, err := hypervisors.ExtractHypervisors(allPages)
	if err != nil {
		panic(err)
	}

	return allNodes, err
}

// 获取project列表
func ListProject(provider *gophercloud.ProviderClient) (result []projects.Project) {
	method := "ListProject"
	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, ":", err)
		return result
	}
	pager := projects.List(client, projects.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		extract, _ := projects.ExtractProjects(page)
		result = extract
		return true, nil
	})
	return result
}

// 获取project列表
func ListPort(provider *gophercloud.ProviderClient) (result []ports.Port) {
	method := "ListPort"
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, ":", err)
		return result
	}
	pager := ports.List(client, ports.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		extract, _ := ports.ExtractPorts(page)
		result = extract
		return true, nil
	})
	return result
}

// 获取flavor列表
func ListFlavor(provider *gophercloud.ProviderClient) (result []flavors.Flavor) {
	method := "ListFlavor"
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, ":", err)
		return result
	}
	pager := flavors.ListDetail(client, flavors.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		extract, _ := flavors.ExtractFlavors(page)
		result = extract
		return true, nil
	})
	return result
}

// 获取VolumesSize
func GetVolumesSizeV2(provider *gophercloud.ProviderClient, Volumes []servers.AttachedVolume) (VolumesSize int) {
	method := "GetVolumesSize"
	client, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, ":", err)
		return
	}
	for _, item := range Volumes {
		if vol, err := volumes.Get(client, item.ID).Extract(); err != nil {
			logging.Error("获取块存储信息出错", err)
		} else {
			VolumesSize += vol.Size
		}
	}
	return
}

// 获取VolumesSizev2
func GetVolumesSize(provider *gophercloud.ProviderClient, Volumes []servers.AttachedVolume) (VolumesSize int) {
	method := "GetVolumesSize"
	client, err := openstack.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		logging.Error(method, ":", err)
		return
	}
	for _, item := range Volumes {
		if vol, err := volumes.Get(client, item.ID).Extract(); err != nil {
			logging.Error("获取块存储信息出错", err)
		} else {
			VolumesSize += vol.Size
		}
	}
	return
}

// 获取vncConsole
func (o *OpenStack) GetVncConsole(ID string) (Consoleurl string, err error) {
	//获取认证
	provider, err := o.Auth()
	if err != nil {
		logging.Error(o.ClusterName, "认证失败:", err)
		return
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	Consoleurl, err = servers.GetVnc(client, ID)
	return Consoleurl, err
}

// openstack 获取认证
func (o *OpenStack) Auth() (*gophercloud.ProviderClient, error) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: o.IdentityEndpoint,
		Username:         o.Username,
		Password:         o.Password,
		DomainName:       o.DomainName,
		TenantName:       o.ProjectName,
		AllowReauth:      true,
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		logging.Error("AuthenticatedClient : %v", err)
	}
	logging.Info("get auth of", o.ClusterName, "...................")
	return provider, err
}

//https://blog.csdn.net/abebrcj842175/article/details/101751887
