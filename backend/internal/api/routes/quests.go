package api_routes

import (
	"git-quest-be/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var questRoutes = []Routes{
	{
		Path:    "",
		Method:  http.MethodGet,
		Handler: getQuests,
	},
	{
		Path:    "/:title",
		Method:  http.MethodGet,
		Handler: getQuest,
	},
	{
		Path:    "/:title",
		Method:  http.MethodPost,
		Handler: startQuest,
	},
	{
		Path:    "/:title/prompt",
		Method:  http.MethodGet,
		Handler: getQuestPrompt,
	},
	{
		Path:    "/:title/validate",
		Method:  http.MethodPost,
		Handler: validateQuest,
	},
	{
		Path:    "/:title",
		Method:  http.MethodDelete,
		Handler: teardownQuest,
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
	title := c.Param("title")

	q, err := services.GetQuest(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, q.GetMeta())
}

func startQuest(c *gin.Context) {
	title := c.Param("title")

	q, err := services.GetQuest(title)
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

func getQuestPrompt(c *gin.Context) {
	title := c.Param("title")

	q, err := services.GetQuest(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"prompt": q.Prompt(),
	})
}

func validateQuest(c *gin.Context) {
	title := c.Param("title")

	q, err := services.GetQuest(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	pass, err := q.Validate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if !pass {
		c.JSON(http.StatusOK, gin.H{
			"message": "quest failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "quest passed",
	})
}

func teardownQuest(c *gin.Context) {
	title := c.Param("title")

	q, err := services.GetQuest(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	err = q.Teardown()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
