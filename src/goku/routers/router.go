package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goku/component"
	"goku/routers/api/v1"
	"goku/utils/logging"
	"goku/utils/options"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(options.Conf.Http.RunMode)
	r.Use(gin.Logger(), gin.Recovery())
	login := v1.LoginResource{}
	r.POST("/api/v1/login", login.Login)
	r.Use(component.JWTAuth())
	apiV1 := r.Group("/api/v1")
	//注册路由
	{
		v1.VMRouter()
		v1.NetWorkRouter()
	}

	for key, instance := range v1.APIs {
		for uri, _func := range instance {
			_value, ok := _func.(func(*gin.Context))
			if !ok {
				panic("invalid api type")
			}
			switch uri.GetModel() {
			case "GET":
				apiV1.GET(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			case "POST":
				apiV1.POST(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			case "DELETE":
				apiV1.DELETE(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			case "PUT":
				apiV1.PUT(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			case "OPTIONS":
				apiV1.OPTIONS(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			case "PATCH":
				apiV1.PATCH(fmt.Sprintf("%s%s", key, uri.GetUri()), _value)
			default:
				logging.Info("no match http request method")
			}
		}
	}
	return r
}
