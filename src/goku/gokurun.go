package main

import (
	"goku/component"
	"goku/pkg/netbox"
	op "goku/pkg/openstack"
	"goku/routers"
	"goku/utils/logging"
	"goku/utils/options"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron"
)

// @title GCloud API
// @description  This is a sample server Petstore server.
// @version 1.0
// @BasePath /api/v1
// @HOST 192.168.12.18:18080
func main() {
	//初始化配置文件
	conf := options.InitConfig()
	//初始化log
	logging.ConfigInit()
	s := &http.Server{
		Addr:           conf.Http.Addr,
		Handler:        routers.InitRouter(),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	netbox.InitFunc()
	//op.Do(component.ClusterInfo)
	//op.DeleteVMFromNetbox(component.ClusterInfo)

	c := cron.New()
	_ = c.AddFunc("0 0 6 * * ?", func() {
		go op.Do(component.ClusterInfo)
	})
	if options.Conf.Netbox.Onoff {
		_ = c.AddFunc("0 0 8 * * ?", func() {
			go op.DeleteVMFromNetbox(component.ClusterInfo)
		})
	}
	c.Start()

	if err := s.ListenAndServe(); err != nil {
		logging.Error(errors.WithStack(err))
	}
}
