package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", Users)
	http.HandleFunc("/user", User)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)

	fmt.Println("started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}