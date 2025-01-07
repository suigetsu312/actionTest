package test
import (
	"net/http"
	"net/http/httptest"
	"example/gintest/db"
	"example/gintest/routes"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// 使用內存數據庫進行測試
	db.InitDB(":memory:")
	routes.RegisterPostRoutes(router)
	return router
}

func TestGetAllPosts(t *testing.T) {
	r := setupTestRouter()

	// 插入測試數據
	db.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", "Test Post", "Test Content")

	req, _ := http.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Post")
}
