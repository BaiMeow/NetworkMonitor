package controller

import (
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/gin-gonic/gin"
)

func BGPAnalysisBetweenness(c *gin.Context) {
	asn := c.Query("asn")
	name := c.Param("name")
	betweenness := graph.GetBgpBetweenness(name)
	if asn == "" {
		c.JSON(200, Resp{
			Code: 0,
			Msg:  "ok",
			Data: betweenness,
		})
		return
	}
	asnUint, err := strconv.ParseUint(asn, 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	if _, ok := betweenness[uint32(asnUint)]; ok {
		c.JSON(200, Resp{
			Code: 0,
			Msg:  "ok",
			Data: betweenness[uint32(asnUint)],
		})
	} else {
		c.JSON(404, RespErrASNNotFound)
	}
}

func BGPAnalysisCloseness(c *gin.Context) {
	name := c.Param("name")
	asn := c.Query("asn")
	closeness := graph.GetBgpCloseness(name)
	if asn == "" {
		c.JSON(200, Resp{
			Code: 0,
			Msg:  "ok",
			Data: closeness,
		})
		return
	}
	asnUint, err := strconv.ParseUint(asn, 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}

	if _, ok := closeness[uint32(asnUint)]; ok {
		c.JSON(200, Resp{
			Code: 0,
			Msg:  "ok",
			Data: closeness[uint32(asnUint)],
		})
	} else {
		c.JSON(404, RespErrASNNotFound)
	}
}
