package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zekth/go_qmk/api/internal/controllers"
)

// MakeRoutes Generate the base routes of the API
func MakeRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.Static("/ui", "./ui")
}
