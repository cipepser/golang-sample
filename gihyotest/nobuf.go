package main

import(
	"time"
)

func main() {
	ch := make(chan string)
	
	go func() {
		time.Sleep(time.Second)
		ch <- "a"
	} ()
	<-ch
}