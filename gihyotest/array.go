package main

import(
	"fmt"
	// 型を調べる
	// "reflect"
)

// 値渡し
// func fn(arr [4]string)  {
// 	arr[0] = "x"
// 	fmt.Println(arr)
// }

func main() {
	// define - 1
	// var arr [4]string
	// arr[0] = "a"
	// arr[1] = "b"
	// arr[2] = "c"
	// arr[3] = "d"
	
	// define - 2
	// arr := [4]string{"a", "b", "c", "d"}

	// define - 3
	// arr := [...]string{"a", "b", "c", "d"}
	// fmt.Println(arr[1])
	// fmt.Println(reflect.TypeOf(arr)) 	// 型を調べる→[4]stringとなるので[...]だと初期化時に必要な長さを指定していることがわかる
	
	// 値渡し
	// arr := [4]string{"a", "b", "c", "d"}
	// fn(arr)
	// fmt.Println(arr)
	
}