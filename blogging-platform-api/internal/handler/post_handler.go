package handler

import(
	"blogging-platform-api/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
    Posts []models.Blog
}

// NewPostHandler initializes the handler with an empty slice database
func NewPostHandler() *PostHandler {
    return &PostHandler{
        Posts: []models.Blog{},
    }
}

func (h *PostHandler) GetAllPosts(ctx *gin.Context) {
    term := ctx.Query("term")
    if term == "" {
        ctx.JSON(http.StatusOK, h.Posts) // uses h.Posts instead of the global slice
        return
    }

    filteredPosts := []models.Blog{}
    for i, j := range h.Posts {
        if strings.Contains(strings.ToLower(j.Title), strings.ToLower(term)) || 
           strings.Contains(strings.ToLower(j.Content), strings.ToLower(term)) || 
           strings.Contains(strings.ToLower(j.Category), strings.ToLower(term)) {
            filteredPosts = append(filteredPosts, h.Posts[i])
        }
    }
    ctx.JSON(http.StatusOK, filteredPosts)
}

func(h *PostHandler) CreatePost(ctx *gin.Context){
	inputs := models.CreateBlogInput{}
	
	if err := ctx.ShouldBindJSON(&inputs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// later implemented to replace len(posts) for better increment system
	// maxID :=0
	blogs := models.Blog{
		ID: len(h.Posts)+1,
		Title: inputs.Title,
		Content: inputs.Content,
		Category: inputs.Category,
		Tags: inputs.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	h.Posts = append(h.Posts,blogs)
	ctx.JSON(http.StatusCreated, blogs)
}


func (h *PostHandler) GetPostByID(ctx *gin.Context){
	id,err:=strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID format"})
		return 
	}
	
	for _,t:=range h.Posts{
		if t.ID==id{
			ctx.JSON(http.StatusOK, t)
			return
		}
	}
	
	ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
}

func (h *PostHandler) DeletePost(ctx *gin.Context){
	id,err:=strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error": "Invalid ID format"})
		return 
	}
	
	for i,t:=range h.Posts{
		if t.ID==id{
			h.Posts = append(h.Posts[:i], h.Posts[i+1:]...)
			ctx.Status(http.StatusNoContent)
			return
		}
	}
	
	ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
}

func (h *PostHandler) UpdatePost(ctx *gin.Context){
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

	for i,t:=range h.Posts{
		if t.ID==id{
			h.Posts[i].Title=inputs.Title
			h.Posts[i].Category=inputs.Category
			h.Posts[i].Tags=inputs.Tags
			h.Posts[i].Content=inputs.Content
			h.Posts[i].UpdatedAt=time.Now()
			ctx.JSON(http.StatusOK,h.Posts[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound,gin.H{"error": "Blog post not found"})
}

