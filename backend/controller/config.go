package controller

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Config(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"refresh_interval": conf.Interval,
	})
}
