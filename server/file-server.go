package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-----> Starting HTTP server...")

	//authenticator := auth.NewBasicAuthenticator("secret.com", Secret)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.FileServer(http.Dir("./data")).ServeHTTP(res, req)
	})

	http.ListenAndServe(":8080", nil)
}
