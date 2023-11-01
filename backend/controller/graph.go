package controller

import (
	"github.com/BaiMeow/NetworkMonitor/service"
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
	gh := service.GetOSPF(uint32(parseUint))
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: gh,
	})
}

func BGP(c *gin.Context) {
	gh := service.GetBGP()
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: gh,
	})
}

func List(c *gin.Context) {
	ls := service.ListAvailable()
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: ls,
	})
}
