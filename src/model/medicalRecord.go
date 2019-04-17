package model

import (
	"bytes"
	"encoding/gob"
	"time"
	"util"
)

type MedicalRecord struct {
	Addr    string
	Time    time.Time
	Desease string
	Info    string
	User    string
	Doc     string
}

func (record *MedicalRecord) HashCode() string {
	sr, err := record.Serialize()
	util.CheckErr(err)
	s := util.ToHash(sr)
	return s
}

func (record *MedicalRecord) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(record)
	return buf.Bytes(), err
}
