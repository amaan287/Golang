package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restApi/initilizers"
)
func init(){
	initilizers.LoadEnvVariables()
	
}

func main(){
r:= gin.Default()
r.GET("/",func(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"message":"pong",
	})
})
r.Run()

}