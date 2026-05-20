package models

import (
	"time"
)

type Blog struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateBlogInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Category string `json:"category" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}