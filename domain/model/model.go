package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int
	Password  string
	Email     string
	Nama      string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type Admin struct {
	Id        int
	Password  string
	Email     string
	Nama      string
	Roles     string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type Barang struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      string
}
type OrderBarang struct {
	gorm.Model
	NamaBarang string
	qty        int
	Harga      string
	TotalOrder string
	Status     string
	UserId     int
	BarangID   int
	// Barangs    []Barang
}
