package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zekth/go_qmk/api/internal/environment"
)

func Ping(c *gin.Context) {
	e, _ := c.Get("env")
	env := e.(environment.EnvVars)
	c.JSON(200, gin.H{
		"message": "pong",
		"version": env.Version,
	})
}
