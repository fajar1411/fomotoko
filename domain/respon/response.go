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
	NoRekening string `json:"no_rekening" `
	NamaBank   string `json:"nama_bank"`
	IDPlayer   uint
}
