package main

import (
	"fmt"
	"net/http"
)

//변경점

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("go to hello"))
		fmt.Println("접속 했습니다")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello World "))
		fmt.Println("hello로 접속 했습니다")
	})

	http.ListenAndServe(":5001", nil)
}
