package controller

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BGPAnalysisBetweenness(c *gin.Context) {
	asn := c.Query("asn")
	betweenness := graph.GetBgpBetweenness()
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
	asn := c.Query("asn")
	closeness := graph.GetBgpCloseness()
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
