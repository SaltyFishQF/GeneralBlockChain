package view

import (
	"controller"
	"net/http"
	"strconv"
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
}
