package model

import (
	"fmt"
	"os"
	"time"
)

// "id" bigserial PRIMARY KEY,
// "title" varchar(225) NOT NULL,
// "images" varchar(225) NOT NULL,
// "tag" varchar(255),
// "deskripsi" TEXT NOT NULL,
// "status" BOOLEAN NOT NULL DEFAULT (false),
// "created_by" bigint NOT NULL,
// "master_data_id" INT CONSTRAINT blog_fk_master_data_id REFERENCES master_data (id) ON UPDATE CASCADE ON DELETE CASCADE,
// "created_at" timestamptz NOT NULL DEFAULT (now()),
// "updated_at" timestamptz NOT NULL DEFAULT (now())

type Blog struct {
	Id             uint64    `json:"id"`
	Title          string    `json:"title" validate:"required"`
	Status         bool      `json:"status" validate:"required"`
	Images         string    `json:"images"`
	Tag            string    `json:"tag"`
	Deskripsi      string    `json:"deskripsi"`
	Created_by     uint64    `json:"created_by"`
	Master_data_id uint64    `json:"master_data_id"`
	Category       string    `json:"category"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type funcBlog struct {
	Data  string
	Model *Blog
	Cari  *FilterCari
}

func NewBlog(data string, M *Blog, C *FilterCari) *funcBlog {
	return &funcBlog{
		Data:  data,
		Model: M,
		Cari:  C,
	}
}

var (
	columBlog     = "title, status, images, tag, deskripsi, created_by, master_data_id"
	tabelBlog     = "blog"
	selectBlog    = fmt.Sprintf("blog.id, title, blog.status, CONCAT('%s', images), tag, deskripsi, created_by, master_data_id,  master_data.value as category,%s", os.Getenv("BASE_URL"), DateGlobal())
	joinBlogMater = "master_data on blog.master_data_id = master_data.id"
)

func (r *funcBlog) Create() string {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s', %v, '%s', '%s','%s', %d, %d) RETURNING ID",
		tabelBlog,
		columBlog,
		r.Model.Title,
		r.Model.Status,
		r.Model.Images,
		r.Model.Tag,
		r.Model.Deskripsi,
		r.Model.Created_by,
		r.Model.Master_data_id,
	)
	return query
}

// Userfindemail = "email=$1"
func (r *funcBlog) Detail() string {

	var hasil string
	var data = fmt.Sprintf("SELECT %s FROM %s %s", selectBlog, tabelBlog, joinBlogMater)
	switch r.Data {
	case "status":
		hasil = fmt.Sprintf("%s WHERE status=%t", data, r.Cari.StatusMaster)
	case "id":
		hasil = fmt.Sprintf("%s WHERE id=%d", data, r.Cari.Id)
	case "all":
		hasil = fmt.Sprintf("%s ORDER BY id DESC", data)
	}
	return hasil
}
