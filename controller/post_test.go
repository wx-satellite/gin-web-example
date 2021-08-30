package controller

import (
	"encoding/json"
	"fmt"
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
	//engine.Run()

	// 构造请求
	body := `{"community_id":1,"title":"test","content":"it is a test"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	res := Response{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, res.Code, 0)
}

type Dog struct {
}

func (*Dog) Println() {
	fmt.Println(123)
}

type P interface {
	Println()
}

type D struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Dd struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Ds   []D    `json:"ds"`
}

type T1 struct {
}

func (t T1) Run() {
	t.SetUp()
	t.TearDown()
}

func (T1) SetUp() {
	fmt.Println("T1 set up")
}

func (T1) TearDown() {
	fmt.Println("T1 tear down")
}

type T2 struct {
	T1
}

func (T2) SetUp() {
	fmt.Println("T2 set up")
}

func (T2) TearDown() {
	fmt.Println("T2 tear down")
}

func TestCreatePost1(t *testing.T) {
	t1 := T2{}
	t1.Run()
}
