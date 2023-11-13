package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Email    string
	Nama     string
	Bank     []Bank `gorm:"foreignKey:IDPlayer;references:ID"`
	// TopUp    []TopUp `gorm:"foreignKey:IDPlayer;references:ID"`
	Wallet Wallet `gorm:"foreignKey:IDPlayer;references:ID"`
	// Wallet   []Wallet `gorm:"foreignKey:IDPlayer;references:ID"`
}

type Bank struct {
	gorm.Model
	NoRekening   string
	NamaBank     string
	NamaRekening string
	IDPlayer     uint
}
type Wallet struct {
	gorm.Model
	NamaDompet    string
	AccountWallet string
	Saldo         float64
	IDPlayer      uint
}
type TopUp struct {
	gorm.Model
	Saldo          float64
	IDWallet       uint
	OrderId        string
	PaymentMethod  string
	Status         string
	LinkPembayaran string
	VirtualAccount string
	TransactionId  string
	IdUser         uint
}
