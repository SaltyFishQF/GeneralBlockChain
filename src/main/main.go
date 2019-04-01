package main

import (
	"dao"
	"net/http"
	"view"
)

func main() {
	dao.GetLastBlockNum()
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http://localhost:8080/addTransaction?user=xxx&doc=xxx&value=xxx&type=xxx&nonce=xx
	http.HandleFunc("/addTransaction", view.AddTransaction)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
