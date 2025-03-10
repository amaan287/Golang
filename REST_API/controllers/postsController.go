package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/restApi/initilizers"
	"github.com/restApi/models"
)

func CreatePost (c *gin.Context) {
	var body struct{
		Title string
        Description string
	}
c.Bind(&body)
	post:= models.Post{Title:body.Title,Description:body.Description}
	result:=initilizers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		log.Fatal("Error creating the post: ",result.Error)
		return
	}
	c.JSON(200,gin.H{
        "post":post,
})
}

func GetPosts(c *gin.Context){
var posts []models.Post
initilizers.DB.Find(&posts)
c.JSON(200,gin.H{
	"posts":posts,
})
}

func GetPostById(c *gin.Context){
	//Get id from url 
	id:= c.Param("id")
	var post models.Post
	initilizers.DB.First(&post,id)
	c.JSON(200,gin.H{
		"post":post,
	})
}

func UpdatePost(c*gin.Context){
	id:=c.Param("id")
	var body struct{
		Title string
		Description string
	}
	c.Bind(body)
	var post models.Post 
	initilizers.DB.First(&post,id)
	initilizers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Description: body.Description,
	})
c.JSON(200,gin.H{
	"post":post,
})
}