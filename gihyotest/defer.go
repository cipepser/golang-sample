package main
import(
	"fmt"
)

func main() {
	defer fmt.Println("execute this statement thirdly")
	defer fmt.Println("execute this statement secondly")

	fmt.Println("execute this statement firstly")	
}