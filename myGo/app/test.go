package main

import (
	"Lv/app/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userAccount := model.UserAccount{
		ID:       1,
		Type:     1,
		Account:  "user001",
		Password: "1234",
		Name:     "Alice",
	}

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": userAccount,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
