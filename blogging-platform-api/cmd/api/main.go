package main

import (
	"blogging-platform-api/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	
	h:=handler.NewPostHandler()

	r.GET("/posts", h.GetAllPosts)

	r.POST("/posts", h.CreatePost)

	r.GET("/posts/:id", h.GetPostByID)
	
	r.DELETE("/posts/:id", h.DeletePost)
	
	r.PUT("/posts/:id", h.UpdatePost)

	err:=r.Run(":8080")
	if err!=nil{
		log.Fatal("Server failed to start: ", err)
	}
}

