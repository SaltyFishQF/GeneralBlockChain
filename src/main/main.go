package main

import (
	"controller"
	"net/http"
	"view"
)

func main() {
	controller.CheckGenesisBlock()
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http://localhost:8080/addTransaction?user=xxx&doc=xxx&value=xxx&type=xxx&nonce=xx
	http.HandleFunc("/addTransaction", view.AddTransaction)
	http.HandleFunc("/allBlock", view.GetAllBlock)
	http.HandleFunc("/block", view.GetBlockByHash)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
