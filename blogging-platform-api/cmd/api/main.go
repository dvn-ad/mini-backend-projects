package main

import (
	"blogging-platform-api/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	posts:= []models.Blog{}
	
	r.GET("/posts", func(ctx *gin.Context) {
		term:=ctx.Query("term")
		if term==""{
			ctx.JSON(http.StatusOK,posts)
			return
		}
		filteredPosts:=[]models.Blog{}
		for i,j:=range posts{
			if strings.Contains(strings.ToLower(j.Title),strings.ToLower(term)) || 
			strings.Contains(strings.ToLower(j.Content),strings.ToLower(term)) || 
			strings.Contains(strings.ToLower(j.Category),strings.ToLower(term)){
				filteredPosts=append(filteredPosts, posts[i])
			}
		}
		
		ctx.JSON(http.StatusOK, filteredPosts)
	})

	r.POST("/posts",func(ctx *gin.Context) {
		inputs := models.CreateBlogInput{}
		
		if err := ctx.ShouldBindJSON(&inputs); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// later implemented to replace len(posts) for better increment system
		// maxID :=0
		blogs := models.Blog{
			ID: len(posts)+1,
			Title: inputs.Title,
			Content: inputs.Content,
			Category: inputs.Category,
			Tags: inputs.Tags,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		posts = append(posts,blogs)
		ctx.JSON(http.StatusCreated, blogs)
	})

	r.GET("/posts/:id",func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,err)
			return 
		}
		
		for _,t:=range posts{
			if t.ID==id{
				ctx.JSON(http.StatusOK, t)
				return
			}
		}
		
		ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
	})
	
	r.DELETE("/posts/:id",func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error": "Invalid ID format"})
			return 
		}
		
		for i,t:=range posts{
			if t.ID==id{
				posts = append(posts[:i], posts[i+1:]...)
				ctx.Status(http.StatusNoContent)
				return
			}
		}
		
		ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
	})
	
	r.PUT("/posts/:id",func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error": "Invalid ID format"})
			return 
		}
		
		inputs:=models.CreateBlogInput{}
		if err := ctx.ShouldBindJSON(&inputs); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i,t:=range posts{
			if t.ID==id{
				posts[i].Title=inputs.Title
				posts[i].Category=inputs.Category
				posts[i].Tags=inputs.Tags
				posts[i].Content=inputs.Content
				posts[i].UpdatedAt=time.Now()
				ctx.JSON(http.StatusOK,t)
				return
			}
		}
		ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
	})

	err:=r.Run(":8080")
	if err!=nil{
		log.Fatal("Server failed to start: ", err)
	}
}

