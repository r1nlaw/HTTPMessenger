package main

import (
	"HTTPMessenger/messenger/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/messege", handlers.GetMessege)
	http.HandleFunc("/messege/send", handlers.PostMessege)

	http.ListenAndServe(":8085", nil)
}
