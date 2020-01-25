package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zekth/go_qmk/api/internal/controllers"
	"github.com/zekth/go_qmk/api/internal/dependencies"
	"github.com/zekth/go_qmk/api/internal/graphql"
	"log"
)

func prometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	}
}

// MakeRoutes Generate the base routes of the API
func MakeRoutes(r *gin.Engine, dependencies dependencies.Dependencies) {
	// Serves the UI folder on the root
	r.Use(static.Serve("/", static.LocalFile("./ui", false)))
	r.GET("/metrics", prometheusHandler())
	api := r.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
	}
	h, err := graphql.NewHandler("schema.graphql", dependencies)
	if err != nil {
		log.Fatal(err)
	}
	r.POST("/graphql", h.GetHTTPHandler())
}
