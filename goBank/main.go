package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		fmt.Println("Error in main ")
		log.Fatal(err)

	}
	if err := store.init(); err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(":3000", store)
	server.Run()

}
