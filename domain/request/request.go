package request

type RequestUser struct {
	Id            int
	Password      string `json:"password" form:"password" validate:"required,min=5"`
	Email         string `json:"email" form:"email" validate:"required,email"`
	Name          string `json:"nama" form:"nama" validate:"required,min=5"`
	Nama_dompet   string `json:"nama_dompet" form:"nama_dompet"`
	AccountWallet string
	Saldo         float64 `json:"saldo" form:"saldo"`
}
type RequestBank struct {
	Id           int
	NoRekening   string `json:"no_rekening" form:"no_rekening" validate:"required,min=8"`
	NamaBank     string `json:"nama_bank" form:"nama_bank" validate:"required,min=3"`
	NamaRekening string `json:"nama_rekening" form:"nama_rekening" validate:"required,min=3"`
	IDPlayer     uint
}
type RequestWallet struct {
	Id            int
	Nama_dompet   string  `json:"nama_dompet" form:"nama_dompet" validate:"required,min=8"`
	AccountWallet string  `json:"akun_dompet" form:"akun_dompet"`
	Saldo         float64 `json:"saldo" form:"saldo"`
	IDPlayer      uint
}

type RequestTopUp struct {
	Id            int
	Saldo         float64 `json:"amount" form:"amount"`
	OrderId       string
	PaymentMethod string `json:"metode_pembayaran" form:"metode_pembayaran" validate:"required,min=3"`
	Status        string
	IDWallet      uint
	IDPlayer      uint
	UpdateSaldo   float64
}
type RequestUserDanwallet struct {
	Id            int
	Password      string `json:"password" form:"password" validate:"required,min=5"`
	Email         string `json:"email" form:"email" validate:"required,email"`
	Name          string `json:"nama" form:"nama" validate:"required,min=5"`
	Nama_dompet   string `json:"nama_dompet" form:"nama_dompet"`
	AccountWallet string
	Saldo         float64 `json:"saldo" form:"saldo"`
}

type ReqProfile struct {
	IdUser        uint
	Email         string
	Nama          string
	NamaDompet    string
	AccountWallet string
	Saldo         float64
	NoRekening    string
}
