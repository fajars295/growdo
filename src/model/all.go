package model

import "time"

func DateGlobal() string {
	return "created_at, updated_at"
}

type MasterTabel struct {
	Id         uint64    `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type FilterCari struct {
	Cari         string `json:"cari"`
	Id           uint64 `json:"id"`
	Roles        uint64 `json:"roles"`
	StatusBanner bool   `json:"status_banner"`
}
