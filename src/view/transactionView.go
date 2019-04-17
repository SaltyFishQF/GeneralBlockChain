package view

import (
	"algorithm"
	"controller"
	"encoding/hex"
	"fmt"
	"model"
	"net/http"
	"strconv"
	"time"
	"util"
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	txType := r.Form["type"][0]
	from := r.Form["from"][0]
	to := r.Form["to"][0]
	medical := model.MedicalRecord{
		Desease: r.Form["desease"][0],
		Info:    r.Form["info"][0],
		Time:    time.Now(),
		User:    to,
		Doc:     from,
	}
	hexhash, err := hex.DecodeString(medical.HashCode())
	util.CheckErr(err)
	medical.Addr = hex.EncodeToString(algorithm.Base58Encode(hexhash))
	t, _ := strconv.Atoi(txType)
	controller.AddTransaction(int32(t), from, to, medical)
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
