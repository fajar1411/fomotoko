package request

import "time"

type RequestUser struct {
	Id        int
	Role      string
	Password  string `json:"password" form:"password" validate:"required,min=5"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Name      string `json:"nama" form:"nama" validate:"required,min=5"`
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}
type RequestBarang struct {
	Id         int
	NamaBarang string `json:"nama_barang" form:"nama_barang" validate:"required,min=5"`
	Stok       int    `json:"stok" form:"stok" validate:"required"`
	Harga      string `json:"harga" form:"harga" validate:"required,min=5"`
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdateAt   time.Time
}
type RequestOrder struct {
	Id         int
	NamaBarang string
	qty        int `json:"qty" form:"qty" validate:"required"`
	Harga      string
	Status     string
	TotalOrder string
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdateAt   time.Time
}
