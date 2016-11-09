package main

import(
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	
	go func() {
		time.Sleep(time.Second * 3) // 3秒待つ
		fmt.Println("1:" + <-ch)
	} ()
	ch <- "a" // 1つ目のバッファに書き込む
	ch <- "b" // 2つ目のバッファに書き込む
	fmt.Println("2:" + <-ch) // 3秒経つ前にaが読み出される
	ch <- "c" // 3秒経つ前にcが書き込まれる
	ch <- "d" // 3秒経つ前だと2つのバッファを使い切っているのでブロック
	fmt.Println("3:" + <-ch)
	fmt.Println("4:" + <-ch)
}