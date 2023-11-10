package request

type RequestUser struct {
	Id       int
	Password string `json:"password" form:"password" validate:"required,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Name     string `json:"nama" form:"nama" validate:"required,min=5"`
}
type RequestBank struct {
	Id         int
	NoRekening string `json:"no_rekening" form:"no_rekening" validate:"required,min=8"`
	NamaBank   string `json:"nama_bank" form:"nama_bank" validate:"required,min=3"`
	IDPlayer   uint
}
