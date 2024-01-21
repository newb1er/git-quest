package api_routes

import "github.com/gin-gonic/gin"

type Routes struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
