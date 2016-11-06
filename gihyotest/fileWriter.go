package main

import(
	"log"
	"os"
)


func main() {
	// file生成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	// クローズ処理
	defer file.Close()
	
	// 書き込むデータ
	messeage := []byte("hello world\n")
	
	// file書き込み
	_, err = file.Write(messeage)
	if err != nil {
		log.Fatal(err)
	}
	
}
