package main

import (
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/controller"
	"github.com/BaiMeow/NetworkMonitor/db"
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
	log.Println("init config")
	if err := conf.Init(); err != nil {
		log.Fatalf("init config fail:%v", err)
	}

	log.Println("init db")
	skipUptime := false
	if err := db.Init(); err != nil {
		if errors.Is(err, db.ErrDatabaseDisabled) {
			skipUptime = true
		} else {
			log.Fatalf("init db fail:%v", err)
		}
	}

	if os.Getenv("SKIP_UPTIME") == "true" {
		skipUptime = true
	}

	log.Println("init graph")
	if err := graph.Init(); err != nil {
		log.Fatalf("init graph fail:%v", err)
	}

	if !skipUptime {
		log.Println("init uptime")
		uptime.Init()
	}

	log.Println("run web")
	r := gin.Default()
	r.Use(middleware.Cors())
	// ospf
	r.GET("/api/ospf/:asn", controller.OSPF)
	r.GET("/api/ospf/uptime/:routerId/recent", controller.OSPFRecentUptime)

	// bgp
	r.GET("/api/bgp/:name", controller.BGP)
	up := r.Group("/api/bgp/:name/uptime", func(c *gin.Context) {
		if conf.Influxdb.Addr == "" {
			c.AbortWithStatusJSON(403, controller.RespErrNotEnabled)
		}
	})
	up.GET("/:asn/recent", controller.BGPRecentUptime)
	up.GET("/:asn/links", controller.BGPLinks)
	ana := r.Group("/api/bgp/:name/analysis", func(c *gin.Context) {
		if !conf.Analysis {
			c.AbortWithStatusJSON(403, controller.RespErrNotEnabled)
		}
	})
	ana.GET("/betweenness", controller.BGPAnalysisBetweenness)
	ana.GET("/closeness", controller.BGPAnalysisCloseness)

	// others
	r.GET("/api/list", controller.List)
	r.GET("/api/config", controller.Config)

	// static
	r.StaticFS("/assets/", &staticRouter{"/static/assets"})
	r.StaticFileFS("/avatar.png", "/static/avatar.png", http.FS(FS))
	r.StaticFileFS("/", "/", &staticRouter{"/static"})
	if conf.MetadataRedirect != "" {
		r.GET("/monitor-metadata.json", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, conf.MetadataRedirect)
		})
	} else {
		r.StaticFile("/monitor-metadata.json", "./monitor-metadata.json")
	}
	err := r.Run(":" + strconv.Itoa(conf.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
