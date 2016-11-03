package main

import(
	"fmt"
	// "os"
	// "errors"
	// "log"
)

// 関数
func hello()  {
	fmt.Println("hello")
}

// 引数
// func sum(i, j int){
// 	fmt.Println(i + j)
// }

// 戻り値
// func sum(i, j int) int{
// 	return i + j
// }

// 複数の返り値
// func swap(i, j int) (int, int){
// 	return j, i
// }

// error - 2
// func div(i, j int) (int, error) {
// 	if j == 0{
// 		return 0, errors.New("divided by zero")
// 	}
// 	return i / j, nil
// }

// 名前付き戻り値
// func div(i, j int) (result int, err error){
// 	if j == 0 {
// 		err = errors.New("divided by zero")
// 		return
// 	}
// 	result = i / j
// 	return 
// }

// 関数の代入
var sum func(i, j int) = func(i, j int){
	fmt.Println(i + j)
}


func main() {
	// hello()
	// n := sum(1, 2)

	// sum
	// fmt.Println(sum(1, 2))
	
	// swap
	// x, _ := swap(3, 4)
	// fmt.Println(x)
	
	// error - 1
	// osパッケージのimportが必要
	// 
	// file, err := os.Open("hello.go")
	// 
	// if err != nil{
	// 	fmt.Println("can not open file")
	// }
	// fmt.Println(file)
	
	// error - 2, 名前付き戻り値
	// errows, logパッケージが必要
	// 
	// n, err := div(10, 0)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(n)
	
	// 関数リテラル
	// func(i, j int){
	// 	fmt.Println(i + j)
	// }(2, 4)
	
	sum(2, 4)
	
	
}