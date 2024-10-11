package controller

import (
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func BGPRecentUptime(c *gin.Context) {
	u64asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	asn := uint32(u64asn)
	ups, err := uptime.Last10TickerRecord(asn)
	if err != nil {
		c.JSON(500, RespInternalError)
		return
	}
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: ups,
	})
}

func BGPLinks(c *gin.Context) {
	u64asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	window, err := time.ParseDuration(c.Query("window"))
	if err != nil {
		c.JSON(400, RespErrParamInvalid)
		return
	}
	t, err := time.ParseDuration(c.Query("time"))
	if err != nil {
		c.JSON(400, RespErrParamInvalid)
		return
	}
	links, err := uptime.Links(uint32(u64asn), window, t)
	if err != nil {
		c.JSON(500, RespInternalError)
		return
	}
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: links,
	})
}

func OSPFRecentUptime(c *gin.Context) {
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "not implemented",
	})
}
