package http

import "http/mydb"

var db = mydb.DB

func addTransaction() {
	sql := "insert into tbl_transaction values(?,?)"
	db.Exec(sql)
}
