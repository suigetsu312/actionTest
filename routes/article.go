package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example/gintest/db"
	"example/gintest/models"
)


// RegisterPostRoutes 註冊文章相關的路由
func RegisterPostRoutes(r *gin.Engine) {
	r.GET("/posts", getAllPosts)
	r.POST("/posts", createPost)
}

// getAllPosts 獲取所有文章
func getAllPosts(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, content, created_at FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan post"})
			return
		}
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

// createPost 新增文章
func createPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

