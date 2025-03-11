package main

import (
	"log"
	"os"
)

type Page struct{
	Title string
	Body []byte
}
func(p *Page) save() error{
	filename:=p.Title + ".txt"
	return os.WriteFile(filename,p.Body,0600)
}

func loadPage(title string) *Page{
	filename:= title + ".txt"
	body, err := os.ReadFile(filename)
if err!=nil{
	log.Fatal("Error reading file: ",err)
}
	return &Page{Title:title,Body:body}
}
func main(){
p1:=&Page{Title: "testPage",Body: []byte("This is a sample page.")}
p1.save()
p2:= loadPage("testPage")

log.Println(string(p2.Body))

}