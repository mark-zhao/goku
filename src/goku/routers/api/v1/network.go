package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	op "goku/pkg/openstack"
	"goku/utils/logging"
	"goku/utils/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//注册路由
func NetWorkRouter() {
	APIs["/networks"] = map[UriInterface]interface{}{
		NewUri("GET","/PortList"):(&InstanceResource{}).PortList,
	}
}

type networkResource struct {
}

//返回数据
type  ResponsePortsData struct {
	Ports []op.Port `json:"ports" bson:"ports"`
}

//获取port信息
func (*InstanceResource) PortList(c *gin.Context)  {
	method := "PortList"
	if _, ok := tools.FunAuth(c,method); !ok {
		resp.Render(c, 200, nil, errors.New("没有权限访问"))
		return
	}
	//获取参数
	IP := c.DefaultQuery("ip","")
	logging.Debug(IP)
	//创建mongodb 链接
	cli, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logging.Error("connect mongo error")
		return
	}
	defer cli.Close()

	//获取port
	var port op.Port
	var ports []op.Port
	client := cli.DB("openstack").C("port")
	cErr := client.Find(bson.M{"ip": IP}).One(&port)
	if cErr == nil {
		ports = append(ports,port)
	}
	logging.Debug(ports)
	resp.Render(c,200,ResponsePortsData{ports},nil)
	return
}