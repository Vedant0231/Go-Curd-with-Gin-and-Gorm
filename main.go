package main

import (
	src "go_gin/src"

	"github.com/gin-gonic/gin"
)

func main() {
	src.Initalmigration()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create/user", src.Createuser)
	r.GET("/get/users", src.Getusers)
	r.GET("/get/user/:id", src.Getuser)
	r.PUT("/update/user/:id", src.Updateuser)
	r.DELETE("/delete/user/:id", src.Deleteuser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
