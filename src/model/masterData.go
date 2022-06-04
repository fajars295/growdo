package model

import (
	"fmt"
	"time"
)

//     "id" bigserial PRIMARY KEY,
//     "type" VARCHAR NOT NULL,
//     "value" VARCHAR NOT NULL,
//     "status" BOOLEAN NOT NULL DEFAULT (false),
//     "created_at" timestamptz NOT NULL DEFAULT (now()),
//     "updated_at" timestamptz NOT NULL DEFAULT (now())

type MasterData struct {
	Id         uint64    `json:"id"`
	Type       string    `json:"type" validate:"required"`
	Status     bool      `json:"status" validate:"required"`
	Value      string    `json:"value"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type funcMasterData struct {
	Data  string
	Model *MasterData
	Cari  *FilterCari
}

func NewMasterData(data string, M *MasterData, C *FilterCari) *funcMasterData {
	return &funcMasterData{
		Data:  data,
		Model: M,
		Cari:  C,
	}
}

var (
	columMasterData  = "type, status, value"
	tabelMasterData  = "master_data"
	selectMasterData = fmt.Sprintf("id, type, status, value,%s", DateGlobal())
)

func (r *funcMasterData) Create() string {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s', %v, '%s') RETURNING ID",
		tabelMasterData,
		columMasterData,
		r.Model.Type,
		r.Model.Status,
		r.Model.Value,
	)
	return query
}

// Userfindemail = "email=$1"
func (r *funcMasterData) Detail() string {

	var hasil string
	var data = fmt.Sprintf("SELECT %s FROM %s", selectMasterData, tabelMasterData)
	switch r.Data {
	case "status":
		hasil = fmt.Sprintf("%s WHERE status=%t", data, r.Cari.StatusMaster)
	case "type":
		hasil = fmt.Sprintf("%s WHERE type='%s' AND status=%t", data, r.Cari.Type, true)
	case "id":
		hasil = fmt.Sprintf("%s WHERE id=%d", data, r.Cari.Id)
	case "all":
		hasil = fmt.Sprintf("%s ORDER BY id DESC", data)
	}
	return hasil
}
