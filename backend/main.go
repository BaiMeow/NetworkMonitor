package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
    "errors"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/controller"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/middleware"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/BaiMeow/NetworkMonitor/trace"

	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
	"github.com/golang-jwt/jwt/v5"
)

type staticRouter struct{ Base string }
func (r *staticRouter) Open(name string) (http.File, error) {
	return http.FS(FS).Open(path.Join(r.Base, name))
}

// ===== JWT 登录相关 =====
var jwtSecret = []byte(getenv("JWT_SECRET", "dev-secret-change-me"))
var demoUser = getenv("LOGIN_USER", "admin")
var demoPass = getenv("LOGIN_PASS", "changeme")

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" { return v }
	return def
}
func makeToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) < 8 || auth[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		tokenStr := auth[7:]
		t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !t.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}

func main() {
	log.Println("init config")
	if err := conf.Init(); err != nil { log.Fatalf("init config fail:%v", err) }

	log.Println("init tracer")
	if err := trace.Init(); err != nil { log.Fatalf("init tracer fail:%v", err) }

	log.Println("init db")
	skipUptime := false
	if err := db.Init(); err != nil {
		if errors.Is(err, db.ErrDatabaseDisabled) {
			skipUptime = true
		} else {
			log.Fatalf("init db fail:%v", err)
		}
	}
	if !db.Enabled || os.Getenv("SKIP_UPTIME") == "true" { skipUptime = true }

	log.Println("init graph")
	if err := graph.Init(); err != nil { log.Fatalf("init graph fail:%v", err) }

	if !skipUptime {
		log.Println("init uptime")
		uptime.Init()
	}
	controller.Init()

	log.Println("run web")
	r := gin.Default()
	r.Use(middleware.Cors())

	// ===== 登录：POST /api/login =====
	r.POST("/api/login", func(c *gin.Context) {
		var req struct{ Username, Password string }
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"bad request"}); return
		}
		if req.Username != demoUser || req.Password != demoPass {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"bad credentials"}); return
		}
		tok, err := makeToken(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"token error"}); return
		}
		c.JSON(http.StatusOK, gin.H{"token": tok})
	})

	// ===== 受保护：GET /api/ping?host=IP =====
	r.GET("/api/ping", authMiddleware(), func(c *gin.Context) {
		host := c.Query("host")
		if host == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing host"}); return
		}
		p, err := ping.NewPinger(host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
		}
		p.SetPrivileged(false)
		p.Count = 1
		p.Timeout = 1 * time.Second

		start := time.Now()
		err = p.Run()
		stats := p.Statistics()
		alive := err == nil && stats.PacketsRecv > 0

		c.JSON(http.StatusOK, gin.H{
			"alive":  alive,
			"rtt_ms": float64(stats.AvgRtt.Microseconds()) / 1000.0,
			"took":   time.Since(start).Milliseconds(),
		})
	})

	// ===== 其它 API（顶层注册，不要放进任何 handler 里）=====
	// OSPF
	{
		ospf := r.Group("/api/ospf")
		ospf.GET("/:asn", controller.OSPF)
		up := ospf.Group("/:asn/uptime", func(c *gin.Context) {
			if conf.Influxdb.Addr == "" {
				c.AbortWithStatusJSON(403, controller.RespErrNotEnabled)
			}
		})
		up.GET("/:routerId/recent", controller.OSPFRecentUptime)
		up.GET("/:routerId/links", controller.OSPFLinks)
	}
	// BGP
	{
		bgp := r.Group("/api/bgp")
		bgp.GET("/:name", controller.BGP)
		up := bgp.Group("/:name/uptime", func(c *gin.Context) {
			if conf.Influxdb.Addr == "" {
				c.AbortWithStatusJSON(403, controller.RespErrNotEnabled)
			}
		})
		up.GET("/:asn/recent", controller.BGPRecentUptime)
		up.GET("/:asn/links", controller.BGPLinks)
		ana := bgp.Group("/:name/analysis", func(c *gin.Context) {
			if !conf.Analysis {
				c.AbortWithStatusJSON(403, controller.RespErrNotEnabled)
			}
		})
		ana.GET("/betweenness", controller.BGPAnalysisBetweenness)
		ana.GET("/closeness", controller.BGPAnalysisCloseness)
	}
	// others
	r.GET("/api/list", controller.List)
	r.GET("/api/config", controller.Config)
	// sse
	r.GET("/api/update", controller.HeadersMiddleware(), controller.Stream.ServeHTTP(), controller.UpdateEvent)
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

	// ===== 最后启动 =====
	if err := r.Run(":" + strconv.Itoa(conf.Port)); err != nil {
		log.Fatalln(err)
	}
}
