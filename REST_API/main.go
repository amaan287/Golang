package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restApi/controllers"
	"github.com/restApi/initilizers"
)
func init(){
	initilizers.LoadEnvVariables()
	initilizers.ConnectToDb()
	
}

func main(){
r:= gin.Default()
r.POST("/posts",controllers.CreatePost)
r.GET("/posts",controllers.GetPosts)
r.GET("/posts/:id",controllers.GetPostById)
r.Run()

}