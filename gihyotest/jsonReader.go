package main

import(
	"fmt"
	"log"
	"os"
	"encoding/json"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"`
	Age int `json:"age"`
	Address string `json:"address,omitempty"`
	memo string
}

func main() {
	// fileを開く
	file, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// データを読み込む変数
	var person Person
	
	// デコーダ
	decorder := json.NewDecoder(file)
	
	// データの読み込み
	err = decorder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	
	// 表示
	fmt.Println(person)
	
}
