package api_routes

import (
	"git-quest-be/internal/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var questRoutes = []Routes{
	{
		Path:    "",
		Method:  http.MethodGet,
		Handler: getQuests,
	},
	{
		Path:    "/:id",
		Method:  http.MethodGet,
		Handler: getQuest,
	},
	{
		Path:    "/:id/start",
		Method:  http.MethodPost,
		Handler: startQuest,
	},
}

func InitQuestRoutes(r *gin.RouterGroup) {
	for _, route := range questRoutes {
		r.Handle(route.Method, route.Path, route.Handler)
	}
}

func getQuests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"quests": services.GetQuests(),
	})
}

func getQuest(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	q, err := services.GetQuest(num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, q.GetMeta())
}

func startQuest(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	q, err := services.GetQuest(num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	err = q.Setup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "quest started",
		"path":    q.GetPath(),
	})
}
