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

// 可変長引数
// func sum(nums ...int) (result int) {
// 	for _, n := range nums{
// 		result += n
// 	}
// 	return
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

	// 可変長引数
	// fmt.Println(sum(1,2,3,4,5))

	// map - 1
	// var month map[int]string = map[int]string{}
	// month[1] = "January"
	// month[2] = "February"
	// fmt.Println(month)

	// map - 2
	month := map[int]string{
		1: "January",
		2:  "February",
	}
	// jan, ok := month[1]
	// fmt.Println(jan)
	// fmt.Println(ok)
	
	// delete(month, 1)
	// fmt.Println(month)
	
	for key, value := range month {
		fmt.Printf("%d %s\n", key, value)
	}

}