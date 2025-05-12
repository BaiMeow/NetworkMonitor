package controller

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/gin-gonic/gin"
	"net/netip"
	"regexp"
	"strconv"
	"time"
)

var bgpNameRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")

func BGPRecentUptime(c *gin.Context) {
	bgpName := c.Param("name")
	if !bgpNameRegex.MatchString(fmt.Sprintf("bgp-%s", bgpNameRegex)) {
		c.JSON(400, RespErrParamInvalid)
		return
	}
	u64asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	asn := uint32(u64asn)

	ups, err := uptime.Last10BGPTickerRecord(bgpName, asn)
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
	u32asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
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
	bgpName := c.Param("name")
	if !bgpNameRegex.MatchString(bgpName) {
		c.JSON(400, RespErrParamInvalid)
		return
	}
	links, err := uptime.BGPLinks(bgpName, uint32(u32asn), window, t)
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
	u64asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	asn := uint32(u64asn)
	routerId := c.Param("routerId")
	addr, err := netip.ParseAddr(routerId)
	if err != nil || !addr.Is4() {
		c.JSON(400, RespErrParamInvalid)
		return
	}
	routerId = addr.String()

	ups, err := uptime.Last10OSPFTickerRecord(asn, routerId)
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

func OSPFLinks(c *gin.Context) {
	u64asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	routerId := c.Param("routerId")
	_, err = netip.ParseAddr(routerId)
	if err != nil {
		c.JSON(400, RespErrParamInvalid)
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
	links, err := uptime.OSPFLinks(uint32(u64asn), routerId, window, t)
	if err != nil || links == nil {
		c.JSON(500, RespInternalError)
		return
	}
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: gin.H{
			"out": links[0],
			"in":  links[1],
		},
	})
}
