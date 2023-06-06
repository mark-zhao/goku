package v1

import (
	"errors"
	"goku/component"
	op "goku/pkg/openstack"
	"goku/utils/logging"
	"goku/utils/tools"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 注册路由
func VMRouter() {
	APIs["/servers"] = map[UriInterface]interface{}{
		//NewUri("GET", "/Search"):(&InstanceResource{}).SearchVM,
		NewUri("GET", "/ServerList"):        (&InstanceResource{}).ServerList,
		NewUri("GET", "/SyncData"):          (&InstanceResource{}).SyncData,
		NewUri("GET", "/GetBill"):           (&InstanceResource{}).GetBill,
		NewUri("GET", "/CreateBill"):        (&InstanceResource{}).CreateBill,
		NewUri("POST", "/CreateVNCConsole"): (&InstanceResource{}).CreateVNCConsole,
	}
}

const (
	timeFormart = "2006-01-02 15:04:05"
)

type InstanceResource struct {
}

type ResponseServersData struct {
	Servers  []op.Server     `json:"servers" bson:"servers"`
	Projects []op.ProjectMap `json:"projects" bson:"projects"`
	Clusters []string        `json:"clusters"`
}

type CreateVNCConsoleReq struct {
	ID          string `json:"uuid" form:"uuid"`
	ClusterName string `json:"cluster" form:"cluster"`
}

type ResponseVNCConsoleData struct {
	Url string `json:"url"`
}

// 账单回复信息
type ResponseBillData struct {
	Date    string      `json:"date"`
	Cluster string      `json:"cluster"`
	Aggdata interface{} `json:"aggdata"`
}

// 获取服务器列表
func (*InstanceResource) ServerList(c *gin.Context) {
	method := "ServerList"
	_, ok := tools.FunAuth(c, method)
	if !ok {
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	//获取集群
	var clusters []string
	//获取参数
	projectName := c.DefaultQuery("project", "")
	IP := c.DefaultQuery("ip", "")
	HostName := c.DefaultQuery("host", "")
	clusterName := c.DefaultQuery("cluster", "")
	Day := c.DefaultQuery("date", "")
	name := c.DefaultQuery("name", "")

	//生成集合日期
	CnameTime := time.Now().Format("2006-01-02")
	if Day == "" && clusterName != "" {
		clusterName = clusterName + CnameTime
	} else if clusterName != "" && Day != "" {
		clusterName = clusterName + Day
	}
	//创建mongodb 链接
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()
	//获取项目
	var project op.ProjectMap
	var projects []op.ProjectMap
	client := cli.DB("openstack").C("project")
	iter := client.Find(nil).Limit(150).Iter()
	for iter.Next(&project) {
		projects = append(projects, project)
	}
	if err := iter.Close(); err != nil {
		logging.Error("项目迭代失败")
	}
	if IP == "" && name == "" && clusterName == "" && HostName == "" {
		logging.Info("ip 、名称 、集群和主机名 不能同时为空")
		resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters}, nil)
		return
	}
	//获取instances
	var server op.Server
	var servers []op.Server
	var date string
	if Day == "" {
		date = CnameTime
	} else {
		date = Day
	}
	if name == "" {
		if len(IP) != 0 {

			for _, cName := range clusters {
				client := cli.DB("openstack").C(cName + date)
				cErr := client.Find(bson.M{"ip": IP}).One(&server)
				if cErr == nil {
					servers = append(servers, server)
					break
				}
			}
			resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
			return
		}
		if len(HostName) != 0 {
			for _, cName := range clusters {
				client := cli.DB("openstack").C(cName + date)
				cErr := client.Find(bson.M{"host": HostName}).All(&servers)
				if cErr == nil {
					resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
					return
				} else {
					resp.Render(c, 200, nil, errors.New("获取数据失败"))
					return
				}
			}
		}
		if len(clusterName) != 0 {
			client := cli.DB("openstack").C(clusterName)
			if len(projectName) != 0 {
				_ = client.Find(bson.M{"project": projectName}).All(&servers)
				resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
				return
			} else {
				_ = client.Find(nil).All(&servers)
				resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
				return
			}
		}
	} else {
		if len(IP) != 0 {
			for _, cName := range clusters {
				client := cli.DB("openstack").C(cName + date)
				cErr := client.Find(bson.M{"ip": IP}).One(&server)
				if cErr == nil {
					servers = append(servers, server)
					break
				}
			}
			resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
			return
		}
		if len(HostName) != 0 {
			for _, cName := range clusters {
				client := cli.DB("openstack").C(cName + date)
				cErr := client.Find(bson.M{"host": HostName}).All(&servers)
				if cErr == nil {
					resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
					return
				} else {
					resp.Render(c, 200, nil, errors.New("获取数据失败"))
					return
				}
			}
		}
		if len(clusterName) != 0 {
			client := cli.DB("openstack").C(clusterName)
			if len(projectName) != 0 {
				_ = client.Find(bson.M{"project": projectName, "name": bson.M{"$regex": name, "$options": "im"}}).All(&servers)
				resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
				return
			} else {
				_ = client.Find(bson.M{"name": bson.M{"$regex": name, "$options": "im"}}).All(&servers)
				resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: servers}, nil)
				return
			}

		} else {
			var Ss []op.Server
			for _, cName := range clusters {
				client := cli.DB("openstack").C(cName + date)
				cErr := client.Find(bson.M{"name": bson.M{"$regex": name, "$options": "im"}}).All(&servers)
				if cErr == nil {
					Ss = append(Ss, servers...)
				} else {
					logging.Error(cErr)
				}
			}
			resp.Render(c, 200, ResponseServersData{Projects: projects, Clusters: clusters, Servers: Ss}, nil)
			return
		}
	}

}

// 获取账单
func (*InstanceResource) GetBill(c *gin.Context) {
	method := "GetBill"
	if _, ok := tools.FunAuth(c, method); !ok {
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	//获取参数
	clusterName := c.DefaultQuery("cluster", "")
	Day := c.DefaultQuery("date", "")
	//生成集合日期
	CDate := time.Now().Format("2006-01-02")
	var Cname string
	if Day == "" {
		Cname = "bill" + CDate
	} else {
		Cname = "bill" + Day
	}
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()
	client := cli.DB("openstack").C(Cname)
	var Rpdata ResponseBillData
	var Rpdatas []ResponseBillData
	if len(clusterName) != 0 {
		cErr := client.Find(bson.M{"cluster": clusterName}).One(&Rpdata)
		if cErr == nil {
			Rpdatas = append(Rpdatas, Rpdata)
			resp.Render(c, 200, Rpdatas, nil)
			return
		}
	}
	iter := client.Find(nil).Limit(150).Iter()
	for iter.Next(&Rpdata) {
		Rpdatas = append(Rpdatas, Rpdata)
	}
	resp.Render(c, 200, Rpdatas, nil)
	return
}

// 生产账单
func (*InstanceResource) CreateBill(c *gin.Context) {
	method := "CreateBill"
	if _, ok := tools.FunAuth(c, method); !ok {
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	Day := c.DefaultQuery("date", "")
	logging.Debug(Day)
	//生成集合日期
	CDate := time.Now().Format("2006-01-02")
	var Cname string
	if Day == "" {
		Cname = "bill" + CDate
	} else {
		Cname = "bill" + Day
	}
	//创建mongodb 链接
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()
	//生产时间字符串
	//date := time.Now().Format(timeFormart)
	//清空账单
	_, err = cli.DB("openstack").C(Cname).RemoveAll(nil)
	if err != nil {
		logging.Error("清空账单集合失败 Err:", err)
	}
	//获取集群
	for _, clusterDetail := range component.ClusterInfo {
		cluster := clusterDetail.ClusterName
		var date string
		if Day == "" {
			date = CDate
		} else {
			date = Day
		}
		client := cli.DB("openstack").C(cluster + date)
		groupStage := bson.M{"$group": bson.M{
			"_id":     bson.M{"project": "$project"},
			"servers": bson.M{"$sum": 1},
			"vcpus":   bson.M{"$sum": "$flavor.vcpus"},
			"ram":     bson.M{"$sum": "$flavor.ram"},
			"disk":    bson.M{"$sum": "$volumes_size"},
		},
		}
		pipe := client.Pipe([]bson.M{groupStage})
		var result []bson.M
		err := pipe.All(&result)
		if err != nil {
			logging.Error(err)
		}
		logging.Debug(result)
		data := ResponseBillData{date, cluster, result}
		client = cli.DB("openstack").C(Cname)
		cErr := client.Insert(data)
		if cErr != nil {
			logging.Error("insert mongo error:", cErr)
		}
	}
	resp.Render(c, 200, nil, nil)
}

// 获取远程地址
func (*InstanceResource) CreateVNCConsole(c *gin.Context) {
	method := "CreateVNCConsole"
	if _, ok := tools.FunAuth(c, method); !ok {
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	//ClusterName := c.DefaultQuery("cluster","")
	//ID := c.DefaultQuery("uuid","")
	var req CreateVNCConsoleReq
	if c.Bind(&req) == nil {
		logging.Debug(req)
		if req.ClusterName != "" && req.ID != "" {
			ClusterName := req.ClusterName
			ID := req.ID
			for _, cluster := range component.ClusterInfo {
				if cluster.ClusterName == ClusterName {
					clusterInfo := op.OpenStack{cluster}
					if url, err := clusterInfo.GetVncConsole(ID); err != nil {
						logging.Error("获取vncconsole 失败:", err)
						resp.Render(c, 200, nil, err)
						return
					} else {
						logging.Debug("创建 vnc 成功：", url)
						resp.Render(c, 200, ResponseVNCConsoleData{Url: url}, nil)
						return
					}
				}
			}
		} else {
			logging.Debug("创建 vnc 时参数传输错误")
			resp.Render(c, 200, nil, errors.New("传输参数出错"))
			return
		}
	} else {
		logging.Debug("创建 vnc 时参数传输错误")
		resp.Render(c, 200, nil, errors.New("传输参数出错"))
		return
	}

}

// 同步数据
func (*InstanceResource) SyncData(c *gin.Context) {
	method := "SyncData"
	if _, ok := tools.FunAuth(c, method); !ok {
		logging.Error("没有足够的权限")
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	go op.Do(component.ClusterInfo)
	resp.Render(c, 200, nil, nil)
}
