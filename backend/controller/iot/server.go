package iot

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/gin-gonic/gin"
)

func Serve() {
	e := gin.Default()
	e.GET("/dn11-lantern/:graph/:asn/peer-count", func(c *gin.Context) {
		graphName := c.Param("graph")
		asnStr := c.Param("asn")
		asn, err := strconv.ParseUint(asnStr, 10, 32)
		if err != nil {
			c.String(http.StatusOK, "0")
			return
		}
		gr := graph.GetBGP(graphName)
		if gr == nil {
			c.String(http.StatusOK, "0")
			return
		}
		data, _ := gr.GetData()
		if data == nil {
			c.String(http.StatusOK, "0")
			return
		}
		var countPeer int
		for _, link := range data.Link {
			if link.Src == uint32(asn) || link.Dst == uint32(asn) {
				countPeer++
			}
		}
		c.String(http.StatusOK, strconv.Itoa(countPeer))
	})
	go func() {
		err := e.Run(fmt.Sprintf(":%d", conf.Iot.Port))
		if err != nil {
			slog.Warn("Iot server stopped ", err)
		}
	}()
}
