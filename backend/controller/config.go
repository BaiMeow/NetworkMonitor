package controller

import (
	"net/http"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/gin-gonic/gin"
)

func Config(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"refresh_interval": conf.Interval,
	})
}
