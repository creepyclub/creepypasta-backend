package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	r := gin.Default()
	r.GET("/", MainHandler)
	log.Fatal(r.Run("localhost:9000"))
}
