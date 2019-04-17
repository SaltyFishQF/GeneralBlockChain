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
	from := r.Form["from"][0]
	to := r.Form["to"][0]
	record := r.Form["recAddr"][0]
	userAec := r.Form["aec"][0]
	fromSign := r.Form["sign"][0]
	t, _ := strconv.Atoi(txType)
	controller.AddTransaction(int32(t), from, to, record, userAec, fromSign)
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
