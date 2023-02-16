package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var port string

func init() {
	flag.StringVar(&port, "p", ":8081", "listen port, eg: -p :8081")
}
func main() {
	flag.Parse()
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "default index"})
	})
	r.Static("/img", "./img")
	err := r.Run(port)
	if err != nil {
		return
	}
}
