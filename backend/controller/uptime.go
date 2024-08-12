package controller

import (
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/gin-gonic/gin"
	"strconv"
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
		return
	}
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: ups,
	})
}

func OSPFRecentUptime(c *gin.Context) {
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "not implemented",
	})
}
