package dao

func GetUserTxNum(key string) (uint64, error) {
	var num uint64
	sql := "select tx_number from tbl_user where user_key = ?"
	err := db.QueryRow(sql, key).Scan(&num)
	return num, err
}
