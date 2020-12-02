package main

import (
	"log"
	"net/http"

	"github.com/qihach/go-week-02/handlers"
)

func main() {

	http.HandleFunc("/blogs", handlers.GetBlogs)

	log.Println("try http://localhost:8090/blogs?user_id=123 and check the log")
	http.ListenAndServe(":8090", nil)
}
