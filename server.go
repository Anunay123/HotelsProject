package main

import (
	"HotelsProject/Controller"
	middlewares "HotelsProject/MiddleWares"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)
import "fmt"

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger())

	return r
}

func main() {
	setupLogOutput()
	r := getRouter()


	r.GET("/hotels/:cityID", Controller.GetData())

	fmt.Println("Works")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
