package main

import(
	"fmt"
	"log"
	"io/ioutil"
)


func main() {
	// 書き込み
	hello := []byte("hello world\n")
	err := ioutil.WriteFile("./fileiotuil.txt", hello, 0666)
	if err != nil {
		log.Fatal(err)
	}
	
	// 読み込み
	message, err := ioutil.ReadFile("./fileiotuil.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))
	
}
