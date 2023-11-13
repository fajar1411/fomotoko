package respon

type ResponseLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
type ResponseUser struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
type ResponBank struct {
	NoRekening   string `json:"no_rekening" `
	NamaBank     string `json:"nama_bank"`
	IDPlayer     uint
	NamaRekening string `json:"nama_rekening" `
}
type ResponWallet struct {
	Id          int     `json:"dompet_id"`
	Nama_dompet string  `json:"nama_dompet"`
	Saldo       float64 `json:"saldo"`
	IDPlayer    uint    `json:"player"`
}
type ResponTopUp struct {
	Id            int
	Saldo         float64 `json:"saldo"`
	OrderId       string  `json:"orderid"`
	PaymentMethod string  `json:"metode_pembayaran"`
	Status        string  `json:"status"`
	IDWallet      uint    `json:"idwallet"`
	IDPlayer      uint    `json:"player_id"`
}

type Responuserandwallet struct {
	Id            int
	Password      string  `json:"password"`
	Email         string  `json:"email"`
	Name          string  `json:"nama"`
	Nama_dompet   string  `json:"nama_dompet"`
	AccountWallet string  `json:"akun_wallet"`
	Saldo         float64 `json:"amount"`
}
