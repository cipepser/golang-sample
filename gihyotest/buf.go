package main

import(
	"time"
)

func main() {
	ch := make(chan string, 3)
	
	go func() {
		time.Sleep(time.Second)
		<-ch
	} ()
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
}