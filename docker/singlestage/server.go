package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %s!", request.URL.Path[1:])
	})
	http.ListenAndServe(":8080",nil)
}
