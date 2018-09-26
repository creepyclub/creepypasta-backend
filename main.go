package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/creepypasta-club/creepypasta-backend/models"
	"github.com/creepypasta-club/creepypasta-backend/roach"
)

var (
	db      roach.Roach
	address string
)

func init() {
	address = os.Getenv("CREEPYPASTA_ADDRESS")
	if address == "" {
		address = ":9000"
	}
	host := os.Getenv("CREEPYPASTA_POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("CREEPYPASTA_POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("CREEPYPASTA_POSTGRES_USER")
	if user == "" {
		user = "mycreepypastauser"
	}
	password := os.Getenv("CREEPYPASTA_POSTGRES_PASSWORD")
	if password == "" {
		password = "mycreepypastapassword"
	}
	database := os.Getenv("CREEPYPASTA_POSTGRES_DATABASE")
	if database == "" {
		database = "mycreepypastadb"
	}
	roach, err := roach.New(roach.Config{Host: host, Port: port, User: user, Password: password, Database: database})
	if err != nil {
		log.Panic(err.Error())
	} else {
		db = roach
	}
}

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

	r.Run(address)
	db.Close()
}
