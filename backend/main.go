package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/controller"
	"github.com/BaiMeow/OSPF-monitor/graph"
	"github.com/BaiMeow/OSPF-monitor/middleware"
	"github.com/gin-gonic/gin"
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
	r.Use(middleware.Cors())
	r.StaticFS("/assets/", &staticRouter{"/static/assets"})
	r.GET("/api/graph", controller.GetGraph)
	r.StaticFileFS("/", "/", &staticRouter{"/static"})
	r.StaticFileFS("/bs.jpg", "/static/bs.jpg", http.FS(FS))
	r.Run(":" + strconv.Itoa(conf.Port))
}
