package dao

import (
	"model"
	"util"
)

func SaveMedicalRecord(record model.MedicalRecord) int64 {
	sql := "insert into tbl_medical_record (medical_record_id, add_time, desease, record_info, userMR_key, doctor_key) values (?,?,?,?,?,?)"
	res, err := db.Exec(sql, record.Addr, record.Time, record.Desease, record.Info, record.User, record.Doc)
	util.CheckErr(err)
	i, err := res.LastInsertId()
	util.CheckErr(err)
	return i
}
