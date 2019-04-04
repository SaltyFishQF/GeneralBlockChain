package view

import (
	"controller"
	"fmt"
	"net/http"
	"strconv"
	"util"
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	txType := r.Form["type"][0]
	user := r.Form["user"][0]
	doc := r.Form["doc"][0]
	value := r.Form["value"][0]
	nonce, _ := strconv.ParseInt(r.Form["nonce"][0], 10, 64)
	t, _ := strconv.Atoi(txType)
	controller.AddTransaction(int32(t), user, doc, value, uint64(nonce))
	fmt.Fprintln(w, "success")
}

func GetAllTranactionByChainID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	id := r.Form["id"][0]
	iid, _ := strconv.Atoi(id)
	tx := controller.GetAllTranactionByChainID(uint32(iid))
	fmt.Fprintln(w, util.ParseJson(tx))
}
