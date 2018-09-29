package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/creepypasta-club/creepypasta-backend/models"
	"github.com/creepypasta-club/creepypasta-backend/roach"
)

var (
	db      *sql.DB
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
		db = roach.Db
	}
}

func GetTopics(c *gin.Context) {
	topics, err := models.GetAllTopics(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.PureJSON(http.StatusOK, topics)
}

func AddTopic(c *gin.Context) {
	var topic models.Topic
	err := c.BindJSON(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topicId, err := topic.Save(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": topicId})
}

func GetTopic(c *gin.Context) {
	stringID := c.Params.ByName("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topic, err := models.GetTopicByID(id, db)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "topic not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.PureJSON(http.StatusOK, topic)
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/topics", GetTopics)
	v1.GET("/topic/:id", GetTopic)
	v1.POST("/topics", AddTopic)

	r.Run(address)
	db.Close()
}
