package respon

import "time"

type ResponseUser struct {
	Id        int
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdat"`
	DeletedAt time.Time
	UpdateAt  time.Time
}

type ResponseBarang struct {
	NamaBarang string `json:"nama_barang"`
	Stok       int    `json:"stok"`
	Harga      string `json:"harga"`
}
