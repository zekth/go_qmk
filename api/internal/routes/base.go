package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/static"
	"github.com/zekth/go_qmk/api/internal/controllers"
	"github.com/zekth/go_qmk/api/internal/graphql"
)

// MakeRoutes Generate the base routes of the API
func MakeRoutes(r *gin.Engine) {
  // Serves the UI folder on the root
  r.Use(static.Serve("/", static.LocalFile("./ui", false)))
  api := r.Group("/api")
	{
    api.GET("/ping", controllers.Ping)
  }
  g := r.Group("/graphql")
  {
    g.POST("",graphql.Handler())
  }
}
