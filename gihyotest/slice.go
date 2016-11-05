package main

import(
	// "fmt"
	"log"
)

// 値渡し
// func fnV(arr [4]int)  {
// 	for i, _ := range arr {
// 		arr[i] = 0
// 	}
// 	log.Println(arr)
// }

// 参照渡し
// func fnP(arr *[4]int)  {
// 	for i, _ := range arr {
// 		arr[i] = 0
// 	}
// 	log.Println(arr)
// }

// Slice
// func fnV(arr []int)  {
// 	for i, _ := range arr {
// 		arr[i] = 10
// 	}
// 	log.Println(arr)
// }
// 
// func fnP(a *[]int)  {
// 	arr := *a
// 	for i, _ := range arr {
// 		arr[i] = 20
// 	}
// 	log.Println(arr)
// }

// GC - 1
// func del(a []int, i int) []int{
// 	copy(a[i:], a[i+1:])
// 	a = a[:len(a)-1]
// 	return a
// }

// GC - 2 
var zero int
func del(a []int, i int) []int{
	copy(a[i:], a[i+1:])
	a[len(a) - 1] = zero 
	a = a[:len(a)-1]
	return a
}


func main() {
	// 値渡し、参照渡し
	// arr := [4]int{1, 2, 3, 4}
	// fnV(arr)
	// log.Println(arr)
	// fnP(&arr)
	// log.Println(arr)
	
	// Slice
	// arr := []int{1, 2, 3, 4}
	// fnV(arr)
	// log.Println(arr)
	// fnP(&arr)
	// log.Println(arr)
	
	// 宣言
	// arr := []int{1, 2, 3, 4}
	// log.Println(arr, len(arr))
	
	// make()
	// arr := make([]int, 4)
	// log.Println(arr, len(arr))
	
	// append, cap
	// arr := make([]int, 4)
	// arr = append(arr, 5)
	// log.Println(arr, len(arr), cap(arr))
	// "len()は値の入っている数、cap()は内部配列のサイズ"
	
	
	//  GC - 1
	// a := []int{1, 2, 3, 4, 5}
	// a = del(a, 2)
	// log.Println(a, len(a), cap(a))
	
	//  GC - 2 
	a := []int{1, 2, 3, 4, 5}
	a = del(a, 2)
	log.Println(a, len(a), cap(a))
	
}