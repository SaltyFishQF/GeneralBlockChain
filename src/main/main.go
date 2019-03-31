package main

import (
	http2 "http"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http://localhost:8080/addTransaction?user=xxx&doc=xxx&value=xxx&type=xxx&nonce=xx
	http.HandleFunc("/addTransaction", http2.AddTransaction)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
