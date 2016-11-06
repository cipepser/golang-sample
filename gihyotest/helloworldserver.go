package main

import(
	"fmt"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", Indexhandler)
	http.ListenAndServe(":3000", nil)
}