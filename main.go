package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/creepypasta-club/creepypasta-backend/models"
)

func TopicsHandler(c *gin.Context) {
	topics := []models.Topic{
		models.Topic{ID: 1, Title: "Крипота", Text: "здесь будет текст"},
		models.Topic{ID: 2, Title: "Начинает свою", Text: "здесь будет текст"},
		models.Topic{ID: 3, Title: "Жизнь на go", Text: "здесь будет текст"},
	}
	c.JSON(http.StatusOK, topics)
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/topics", TopicsHandler)

	log.Fatal(r.Run("localhost:9000"))
}
