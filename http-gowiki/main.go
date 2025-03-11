package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
  fmt.Fprintf(w, "Select correct message, Current method is %s", r.Method)
	} else {
		fmt.Fprintf(w, "Hi there, I love %s",r.URL.Path[1:])
	}
}
func main() {
	PORT:=":8080"
http.HandleFunc("/",handler)
log.Println(fmt.Sprintf("server is running on http://localhost%s", PORT))
log.Fatal(http.ListenAndServe(PORT, nil))
}