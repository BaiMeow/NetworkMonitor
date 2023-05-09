package main

import (
	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/controller"
	"github.com/BaiMeow/OSPF-monitor/graph"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

type staticRouter struct {
	Base string
}

func (r *staticRouter) Open(name string) (http.File, error) {
	return http.FS(FS).Open(path.Join(r.Base, name))
}

func main() {
	if err := conf.Init(); err != nil {
		log.Fatalf("init config fail:%v", err)
	}
	if err := graph.Init(); err != nil {
		log.Fatalf("init graph fail:%v", err)
	}

	r := gin.Default()
	r.StaticFS("/assets/", &staticRouter{"/static/assets"})
	r.GET("/api/graph", controller.GetGraph)
	r.StaticFileFS("/", "/", &staticRouter{"/static"})
	r.Run(":80")
}
