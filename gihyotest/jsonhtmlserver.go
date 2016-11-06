package main

import(
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func PersonHandler(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	
	// リクエストのbodyをjsonに変換する
	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		// filenameを{id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		
		// fileにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}
		
		// レスポンスとして{201 CREATED}を送信
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}