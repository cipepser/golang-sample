package main

import (
	"fmt"
	"log"
	"errors"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	a := []int{1, 2, 3}

	// fmt.Println(a[10])

	for i := 0; i < 10; i++ {
		if i >= len(a) {
			panic(errors.New("index out of range"))
		}
		fmt.Println(a[i])
	}

}
