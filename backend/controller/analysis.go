package controller

import (
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/gin-gonic/gin"
)

func BGPAnalysisBetweenness(c *gin.Context) {
	name := c.Param("name")
	betweenness := graph.GetBgpBetweenness(name)
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: betweenness,
	})
	return
}

func BGPAnalysisCloseness(c *gin.Context) {
	name := c.Param("name")
	closeness := graph.GetBgpCloseness(name)
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: closeness,
	})
	return
}

func BGPAnalysisPathBetweenness(c *gin.Context) {
	name := c.Param("name")
	betweenness := graph.GetBgpPathBetweenness(name)
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: betweenness,
	})
	return
}

func OSPFAnalysisBetweenness(c *gin.Context) {
	asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	betweenness := graph.GetOSPFBetweenness(uint32(asn))
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: betweenness,
	})
	return
}

func OSPFAnalysisCloseness(c *gin.Context) {
	asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	closeness := graph.GetOSPFCloseness(uint32(asn))
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: closeness,
	})
	return
}

func OSPFAnalysisPathBetweenness(c *gin.Context) {
	asn, err := strconv.ParseUint(c.Param("asn"), 10, 32)
	if err != nil {
		c.JSON(400, RespErrASNInvalid)
		return
	}
	betweenness := graph.GetOSPFPathBetweenness(uint32(asn))
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "ok",
		Data: betweenness,
	})
	return
}
