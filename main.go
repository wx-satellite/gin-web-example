package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Dog struct {
	Name string
}

func main() {
	r := gin.Default()
	r.Use(func(context *gin.Context) {
		fmt.Println("测试")
	})
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	r.Run()
	//s := make([]int64, 0, 10)
	//fmt.Println(fmt.Sprintf("%v", s))
	//s[0] = 100
	//s = append(s, 10)
	//s = append(s, 11)
	//s = append(s, 12)
	//fmt.Println(fmt.Sprintf("%v", s))

}
