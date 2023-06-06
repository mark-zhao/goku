package netbox

import (
	"goku/utils/logging"
	"goku/utils/options"
	"reflect"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/client/tenancy"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
	"github.com/netbox-community/go-netbox/netbox/models"
)

var NewVM chan *Vm
var DueVM chan *DelVM

type Vm struct {
	Name        string
	ClusterName string
	Vcpus       float64
	Memory      int64
	Disk        int64
	Eth0        string
	Eth1        string
	Eth2        string
	Eth3        string
	Tenant      string
	Uuid        string
	Vmcreated   string
	Host        string
}

type DelVM struct {
	ID  int64
	IPS []string
}

type Netbox struct {
	c         *client.NetBoxAPI
	Vm        *Vm
	VmCluster map[string]int64
	VmTenant  map[string]int64
}

//初始化任务队列
func InitFunc() {
	NewVM = make(chan *Vm, 50)
	for i := 0; i < 1; i++ {
		go func() {
			for {
				SyncVM2Netbox(<-NewVM)
			}
		}()
	}
	DueVM = make(chan *DelVM, 50)
	for i := 0; i < 1; i++ {
		go func() {
			for {
				DeleteVMFromNetbox(<-DueVM)
			}
		}()
	}
}

//获得所有虚拟机生产uuid-id map
func GetVMMap() (VMMap map[string]int64, err error) {
	transport := httptransport.New(options.Conf.Netbox.Host, options.Conf.Netbox.Path, client.DefaultSchemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+options.Conf.Netbox.Token)
	c := client.New(transport, nil)
	// 生产VMMap
	req := virtualization.NewVirtualizationVirtualMachinesListParams()
	var limit int64 = 1000
	var offset int64 = 0
	var vmlist = []*models.VirtualMachineWithConfigContext{}
	req.SetLimit(&limit)
	for {
		req.SetOffset(&offset)
		rs, err := c.Virtualization.VirtualizationVirtualMachinesList(req, nil)
		if err != nil {
			return nil, err
		}
		vmlist = append(vmlist, rs.Payload.Results...)
		if *rs.Payload.Count < offset+limit {
			break
		} else {
			offset = offset + limit
		}
	}
	VMMap = make(map[string]int64)
	//logging.Info(rs.Payload.Next,*rs.Payload.Count)
	for _, v := range vmlist {
		if v.CustomFields != nil {
			if value, ok := v.CustomFields.(map[string]interface{}); ok {
				if UUID, ok := value["uuid"]; ok {
					if uuid, ok := UUID.(string); ok {
						VMMap[uuid] = v.ID
					}
				}
			}
		}
	}
	return
}

//Update Primary IP
func (n *Netbox) UpdateVMPrimaryIP(ipAddressID, vmCluster *int64, vm *models.VirtualMachineWithConfigContext) error {
	virtualreq := virtualization.NewVirtualizationVirtualMachinesUpdateParams()
	virtualdata := new(models.WritableVirtualMachineWithConfigContext)
	//virtualdata.ID =  deviceType.ID
	virtualdata.PrimaryIp4 = ipAddressID
	virtualdata.Name = vm.Name
	virtualdata.Cluster = vmCluster
	tags := []*models.NestedTag{}
	virtualdata.Tags = tags
	virtualreq.SetData(virtualdata)
	virtualreq.SetID(vm.ID)
	_, err := n.c.Virtualization.VirtualizationVirtualMachinesUpdate(virtualreq, nil)
	if err != nil {
		logging.Error("update primary ip Error:", err)
		return err
	}
	return nil
}

//初始化Netbox
func newNetbox(vm *Vm) (netbox *Netbox) {
	defer func() {
		if err := recover(); err != nil { //产生了panic异常
			logging.Error(err)
		}
	}()
	//init netbox client
	transport := httptransport.New(options.Conf.Netbox.Host, options.Conf.Netbox.Path, client.DefaultSchemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+options.Conf.Netbox.Token)
	c := client.New(transport, nil)
	netbox = &Netbox{
		c:  c,
		Vm: vm,
	}
	netbox.getCluster()
	netbox.getTenantId()
	logging.Debug(*netbox)
	return netbox
}

//确认netbox里是否存在虚拟机
func (n *Netbox) checkVm() (bool, *models.VirtualMachineWithConfigContext, error) {
	defer func() {
		if err := recover(); err != nil { //产生了panic异常
			logging.Error(err)
		}
	}()
	// check VM exists on netbox
	req := virtualization.NewVirtualizationVirtualMachinesListParams()
	//req.SetName(&name)
	req.SetName(&n.Vm.Name)

	rs, err := n.c.Virtualization.VirtualizationVirtualMachinesList(req, nil)
	if err != nil {
		logging.Error(err)
		return false, nil, err
	}
	if *rs.Payload.Count == 1 {
		//return true,rs.GetPayload().Results[0], nil
		return true, rs.Payload.Results[0], nil
	}
	return false, nil, nil
}

//确认ip是否存在
func (n *Netbox) checkIp(ip *string) (bool, int64, error) {
	// check IP exists on netbox
	req := ipam.NewIpamIPAddressesListParams()
	req.SetAddress(ip)

	rs, err := n.c.Ipam.IpamIPAddressesList(req, nil)

	if err != nil {
		return false, 10000, err
	}

	if *rs.Payload.Count == 1 {
		return true, rs.Payload.Results[0].ID, nil
	}
	return false, 10000, nil
}

//获取集群
func (n *Netbox) getCluster() {
	req := virtualization.NewVirtualizationClustersListParams()
	ClusterRes, err := n.c.Virtualization.VirtualizationClustersList(req, nil)
	if err != nil {
		logging.Error(err)
	}
	ClusterMap := make(map[string]int64)
	for _, v := range ClusterRes.Payload.Results {
		ClusterMap[*v.Name] = v.ID
	}
	n.VmCluster = ClusterMap
}

//获取项目
func (n *Netbox) getTenantId() {
	// get TenantId
	//req := virtualization.NewVirtualizationInterfacesListParams()
	req := tenancy.NewTenancyTenantsListParams()
	var limit int64 = 100
	req.SetLimit(&limit)
	rs, err := n.c.Tenancy.TenancyTenantsList(req, nil)
	if err != nil {
		logging.Error(err)
	}
	TenantMap := make(map[string]int64)
	for _, v := range rs.Payload.Results {
		if v.CustomFields != nil {
			if value, ok := v.CustomFields.(map[string]interface{}); ok {
				if pID, ok := value["projectid"]; ok {
					if projectID, ok := pID.(string); ok {
						TenantMap[projectID] = v.ID
					}
				}
			}
		}
	}
	n.VmTenant = TenantMap
}

func SyncVM2Netbox(vm *Vm) {
	method := "SyncVM2Netbox"
	logging.Info("begin", method)
	//loop vm check and import
	n := newNetbox(vm)
	if n == nil {
		logging.Error("访问netbox 出错")
		return
	}
	vmreq := virtualization.NewVirtualizationVirtualMachinesCreateParams()
	vmdata := new(models.WritableVirtualMachineWithConfigContext)
	//netbox 是否存在此虚拟机
	if T, res, e := n.checkVm(); e == nil {
		if T && res.CustomFields != nil {
			if value, ok := res.CustomFields.(map[string]interface{}); ok {
				if v, ok := value["uuid"]; ok {
					if uuid, ok := v.(string); ok {
						if uuid == vm.Uuid {
							logging.Error(vm.Name, "虚拟机已经存在")
						}
					}
				}
			}
			newName := vm.Name + vm.Vmcreated
			vmdata.Name = &newName
		} else {
			vmdata.Name = &vm.Name
		}
	} else {
		logging.Error("获取云主机访问netbox 出错")
		return
	}
	if vm.Vcpus > 0 && vm.Memory > 0 && vm.Disk > 0 {
		vmdata.Vcpus = &vm.Vcpus
		vmdata.Memory = &vm.Memory
		vmdata.Disk = &vm.Disk
	}
	var cluster int64 = n.VmCluster[n.Vm.ClusterName]
	//vmdata.Cluster = &cluster
	vmdata.Cluster = &cluster
	if tenant, ok := n.VmTenant[n.Vm.Tenant]; ok {
		vmdata.Tenant = &tenant
	}
	var customfildsMap = make(map[string]string)
	customfildsMap["uuid"] = vm.Uuid
	customfildsMap["host"] = vm.Host
	customfildsMap["VMcreated"] = vm.Vmcreated
	var customfildValue interface{} = customfildsMap
	vmdata.CustomFields = customfildValue
	tags := []*models.NestedTag{}
	vmdata.Tags = tags
	vmdata.Created = strfmt.Date{}
	vmdata.LastUpdated = strfmt.DateTime{}
	vmreq.SetData(vmdata)
	vmRes, err := n.c.Virtualization.VirtualizationVirtualMachinesCreate(vmreq, nil)
	if err != nil {
		logging.Error("Create vm", vm.Name, err)
		return
	} else {
		logging.Info("Create vm", vm.Name, "success")
	}
	//ip 信息整理
	immutableT := reflect.TypeOf(*vm)
	var interfs [4]string = [4]string{"Eth0", "Eth1", "Eth2", "Eth3"}
	for index, value := range interfs {
		var T bool
		var e error
		var ip *string
		var ID int64
		var interfaceName string
		switch index {
		case 0:
			if vm.Eth0 == "" {
				continue
			}
			interfaceName = "eth0"
			T, ID, e = n.checkIp(&vm.Eth0)
			ip = &vm.Eth0
		case 1:
			if vm.Eth1 == "" {
				continue
			}
			interfaceName = "eth1"
			T, ID, e = n.checkIp(&vm.Eth1)
			ip = &vm.Eth1
		case 2:
			if vm.Eth2 == "" {
				continue
			}
			interfaceName = "eth2"
			T, ID, e = n.checkIp(&vm.Eth2)
			ip = &vm.Eth2
		case 3:
			if vm.Eth3 == "" {
				continue
			}
			interfaceName = "eth3"
			T, ID, e = n.checkIp(&vm.Eth3)
			ip = &vm.Eth3
		}
		// create interface on netbox
		interreq := virtualization.NewVirtualizationInterfacesCreateParams()
		interdata := new(models.WritableVMInterface)
		interdata.Name = &interfaceName
		tags := []*models.NestedTag{}
		interdata.Tags = tags
		interdata.TaggedVlans = []int64{}
		interdata.VirtualMachine = &vmRes.Payload.ID
		interreq.SetData(interdata)
		intRes, err := n.c.Virtualization.VirtualizationInterfacesCreate(interreq, nil)
		if err != nil {
			logging.Error("Create interface:", n.Vm.Name, err)
			return
		} else {
			logging.Info("Create interface:", n.Vm.Name, "success")
		}
		if _, ok := immutableT.FieldByName(value); ok {
			if e != nil {
				logging.Error("netbox get ip api 出错：", e)
				return
			} else if T == true {
				//update ip on netbox
				logging.Error("发生严重的错误", n.Vm.Name, ":", *ip, "出现冲突")
				ethReq := ipam.NewIpamIPAddressesUpdateParams()
				ethData := new(models.WritableIPAddress)
				if tenant, ok := n.VmTenant[n.Vm.Tenant]; ok {
					ethData.Tenant = &tenant
				}
				objectType := "virtualization.vminterface"
				ethData.AssignedObjectType = &objectType
				ethData.Address = ip
				ethData.Tags = tags
				ethData.AssignedObjectID = &intRes.Payload.ID
				ethReq.SetData(ethData)
				ethReq.ID = ID
				// Update IP
				ipRes, err := n.c.Ipam.IpamIPAddressesUpdate(ethReq, nil)
				if err != nil {
					logging.Error("Update IP Error:", err)
					return
				} else {
					if interfaceName == "eth0" {
						_ = n.UpdateVMPrimaryIP(&ipRes.Payload.ID, &cluster, vmRes.Payload)
					}
				}
			} else if T == false && e == nil {
				// create ip on netbox
				ethReq := ipam.NewIpamIPAddressesCreateParams()
				ethData := new(models.WritableIPAddress)
				if tenant, ok := n.VmTenant[n.Vm.Tenant]; ok {
					ethData.Tenant = &tenant
				}
				objectType := "virtualization.vminterface"
				ethData.AssignedObjectType = &objectType
				ethData.Address = ip
				ethData.Tags = tags
				ethData.AssignedObjectID = &intRes.Payload.ID
				//Create IP
				ethReq.SetData(ethData)
				ipRes, err := n.c.Ipam.IpamIPAddressesCreate(ethReq, nil)
				if err != nil {
					logging.Error("Create IP Error:", err)
					return
				} else {
					if interfaceName == "eth0" {
						_ = n.UpdateVMPrimaryIP(&ipRes.Payload.ID, &cluster, vmRes.Payload)
					}
				}
			} else {
				logging.Error("ip 已存在与netbox系统")
				return
			}

		}
	}
}

func DeleteVMFromNetbox(vm *DelVM) {
	transport := httptransport.New(options.Conf.Netbox.Host, options.Conf.Netbox.Path, client.DefaultSchemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+options.Conf.Netbox.Token)
	c := client.New(transport, nil)
	req := virtualization.NewVirtualizationVirtualMachinesDeleteParams()
	req.SetID(vm.ID)
	_, err := c.Virtualization.VirtualizationVirtualMachinesDelete(req, nil)
	if err != nil {
		logging.Error(err)
		return
	}
	for _, v := range vm.IPS {
		CreateIPTONetbox(&v)
	}
}

func CreateIPTONetbox(ip *string) {
	logging.Info("开始删除")
	transport := httptransport.New(options.Conf.Netbox.Host, options.Conf.Netbox.Path, client.DefaultSchemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+options.Conf.Netbox.Token)
	c := client.New(transport, nil)
	tags := []*models.NestedTag{}
	ethReq := ipam.NewIpamIPAddressesCreateParams()
	ethData := new(models.WritableIPAddress)
	ethData.Description = "云平台使用"
	ethData.Address = ip
	ethData.Tags = tags
	//Create IP
	ethReq.SetData(ethData)
	_, err := c.Ipam.IpamIPAddressesCreate(ethReq, nil)
	if err != nil {
		logging.Error("Create IP Error:", err)
		return
	}
}
