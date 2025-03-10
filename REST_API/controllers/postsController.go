package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/restApi/initilizers"
	"github.com/restApi/models"
)

func CreatePost (c *gin.Context) {
	post:= models.Post{Title:"HEllo",Description:"Description",CreatedAt:time.Now()}
	result:=initilizers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		log.Fatal("Error creating the post: ",result.Error)
		return
	}
	c.JSON(200,gin.H{
		"message":"pong",
})
}
