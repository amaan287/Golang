package main
import ("fmt" 
"net/http"
	"log"
)
func helloHandler(w http.ResponseWriter,r *http.Request){

}
func main(){
	fileserver:= http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/Hello",helloHandler)
	fmt.Printf("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080",nil), err != nil{
		log.Fatal(err)
	} 
	
}
