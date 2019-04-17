package main

import (
	"controller"
	"net/http"
	"view"
)

func main() {
	//prk = "60c65722e914c9665445f14fcf9cf777c5c51ea62597e345c9ecc794730123f0"
	controller.CheckGenesisBlock()
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http://localhost:8080/addTransaction?user=xxx&doc=xxx&value=xxx&type=xxx&nonce=xx
	http.HandleFunc("/addTransaction", view.AddTransaction)
	http.HandleFunc("/allBlock", view.GetAllBlock)
	http.HandleFunc("/block", view.GetBlockByHash)
	http.HandleFunc("/transaction", view.GetAllTranactionByChainID)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	//prk, pub, _ := algorithm.GetKey()
	//rk := algorithm.PrivateKeyToByte(prk)
	//ub := algorithm.PublicKeyToByte(pub)
	//fmt.Println(hex.EncodeToString(rk), hex.EncodeToString(ub))
}
