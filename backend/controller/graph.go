package controller

import (
	"github.com/BaiMeow/OSPF-monitor/service"
	"github.com/gin-gonic/gin"
)

func GetGraph(c *gin.Context) {
	gh := service.GetGraph()
	c.JSON(200, gh)
}
