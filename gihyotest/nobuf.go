package main

import(
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	
	// データが書き込まれるまで待つ場合
	// go func() {
	// 	time.Sleep(time.Second)
	// 	ch <- "a"
	// } ()
	// fmt.Println(<-ch) // 書き込まれるのを待つ
	
	// 読み出されるまで待つ場合
	go func() {
		// fmt.Println(<-ch) 
		time.Sleep(time.Second)
		fmt.Println(<-ch) 
	} ()
	// time.Sleep(time.Second)
	ch <- "a"
}