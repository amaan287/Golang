package main

import (
	"github.com/restApi/initilizers"
	"github.com/restApi/models"
)
func init (){
	initilizers.LoadEnvVariables()
initilizers.ConnectToDb()
}
func main(){
initilizers.DB.AutoMigrate(&models.Post{})
}