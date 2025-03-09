package initilizers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){

	if envErr:= godotenv.Load(); envErr!=nil{;
	   log.Fatal("Error loading environment variable");
	 };
}