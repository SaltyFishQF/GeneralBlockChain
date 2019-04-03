package view

import (
	"controller"
	"fmt"
	"net/http"
	"util"
)

func GetAllBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	blocks := controller.GetAllBlock()
	res := util.ParseJson(blocks)
	fmt.Fprintln(w, res)
}

func GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	block := controller.GetBlockByHash(r.Form["hash"][0])
	res := util.ParseJson(block)
	fmt.Fprintln(w, res)
}
