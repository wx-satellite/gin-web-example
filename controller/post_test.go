package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	url := "/api/v1/post"
	engine.POST(url, CreatePost)

	// 构造请求
	body := `{"community_id":1,"title":"test","content":"it is a test"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	res := Response{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, res.Code, 0)
}
