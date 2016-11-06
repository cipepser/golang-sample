package main

import(
	"fmt"
)

// Type Assertion
// func Print(value interface{} ) {
// 	s, ok := value.(string)
// 	if ok {
// 		fmt.Printf("value is string: %s\n", s)
// 	} else {
// 		fmt.Printf("value is not string\n")
// 	}
// }

// Type Switch
type Stringer interface {
	String() string
}

func Print(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is Stringer: %s\n", v)
	}
}


func main() {
	// キャスト
	// var i uint8 = 3
	// var j uint32 = uint32(i)
	// fmt.Println(j)
	// 
	// var s string = "abc"
	// var b []byte = []byte(s)
	// fmt.Println(b)
	
	// a := int("a") // エラー
	// a := int("1") // エラー
	
	// Type Assertion, Type Switch
	Print("abc")
	Print(10)
	Print()
}
