package controller

import (
	"github.com/BaiMeow/NetworkMonitor/service/graph"
	"github.com/gin-gonic/gin"
	"strconv"
)

func OSPF(c *gin.Context) {
	asn := c.Param("asn")
	parseUint, err := strconv.ParseUint(asn, 10, 32)
	if err != nil {
		c.JSON(400, Resp{
			Code: -1,
			Msg:  "invalid ASN",
		})
		return
	}
	gh, t := graph.GetOSPF(uint32(parseUint))
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: gin.H{
			"updated_at": t,
			"graph":      gh,
		},
	})
}

func BGP(c *gin.Context) {
	name := c.Param("name")
	gh, t := graph.GetBGP(name)
	if gh == nil {
		c.JSON(404, RespErrBGPGraphNotFound)
		return
	}
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: gin.H{
			"as":         gh.AS,
			"link":       gh.Link,
			"updated_at": t,
		},
	})
}

func List(c *gin.Context) {
	ls := graph.ListAvailable()
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: ls,
	})
}
