package api_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var indexRoutes = []Routes{
	{
		Path:    "/",
		Method:  http.MethodGet,
		Handler: healthCheck,
	},
}

func InitIndexRoutes(r *gin.RouterGroup) {
	for _, route := range indexRoutes {
		r.Handle(route.Method, route.Path, route.Handler)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
