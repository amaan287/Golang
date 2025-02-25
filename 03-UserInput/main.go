package main
import ("fmt"
"bufio"
	"os"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")
	
	//comma ok syntax // error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println( "Your rating is: ", input)
fmt.Printf("Type of this rating is %T", input)
}
