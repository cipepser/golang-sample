package main

import(
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()


	message := make([]byte, 12)
	
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(string(message))
}
