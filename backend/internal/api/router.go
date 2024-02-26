package api

import (
	api_routes "git-quest-be/internal/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Path    string
	Action  string
	Handler func() interface{}
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	api_routes.InitIndexRoutes(r.Group("/"))
	api_routes.InitWsRoutes(r.Group("/ws"))
	api_routes.InitQuestRoutes(r.Group("/quests"))

	return r
}
