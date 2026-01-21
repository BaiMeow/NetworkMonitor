package controller

import (
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
