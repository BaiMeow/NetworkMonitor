package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/controller"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/middleware"
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
	r.GET("/api/ospf/:asn", controller.OSPF)
	r.GET("/api/bgp", controller.BGP)
	r.StaticFileFS("/", "/", &staticRouter{"/static"})
	r.StaticFileFS("/avatar.png", "/static/avatar.png", http.FS(FS))
	r.Run(":" + strconv.Itoa(conf.Port))
}
