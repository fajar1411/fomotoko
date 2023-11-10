package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Email    string
	Nama     string
	Bank     []Bank  `gorm:"foreignKey:IDPlayer;references:ID"`
	TopUp    []TopUp `gorm:"foreignKey:IDPlayer;references:ID"`
	Wallet   Wallet  `gorm:"foreignKey:IDPlayer;references:ID"`
	// Wallet   []Wallet `gorm:"foreignKey:IDPlayer;references:ID"`
}

type Bank struct {
	gorm.Model
	NoRekening string
	NamaBank   string
	IDPlayer   uint
	TopUp      []TopUp `gorm:"foreignKey:IDBank;references:ID"`
}
type Wallet struct {
	gorm.Model
	NamaDompet string
	Saldo      float64
	IDPlayer   uint
}
type TopUp struct {
	gorm.Model
	Saldo      float64
	IDBank     uint
	IDWallet   uint
	IDPlayer   uint
	NamaDompet string
	// Wallet   []Wallet `gorm:"foreignKey:IDWallet;references:ID"`
}
