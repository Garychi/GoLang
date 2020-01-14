package router

import (
	"Lv/app/handler"
	"github.com/gin-gonic/gin"
)

func LoadAPIRouter(r *gin.Engine){
	authorization := r.Group("/")

	authorization.GET("/api/allUser", handler.FetchAllUser)
	authorization.GET("/api/user", handler.FetchUser)
}

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}