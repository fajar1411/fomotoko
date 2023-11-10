package query

import (
	"test/domain/model"
	"test/domain/request"
	"test/domain/respon"
)

func RequuserToModel(data request.RequestUser) model.User {
	return model.User{

		Email:    data.Email,
		Nama:     data.Name,
		Password: data.Password,
	}
}
func ModelusertoReq(data *model.User) request.RequestUser {
	return request.RequestUser{
		Email:    data.Email,
		Name:     data.Nama,
		Password: data.Password,
		Id:       int(data.ID),
	}
}

func ReqtoRepon(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Id:       data.Id,
		Email:    data.Email,
		Name:     data.Name,
		Password: data.Password,
	}
}

func ReqtoResponLogin(data request.RequestUser, token string) respon.ResponseLogin {
	return respon.ResponseLogin{

		Email:    data.Email,
		Name:     data.Name,
		Token:    token,
		Password: data.Password,
	}
}

func ListModeluserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModelusertoReq(&val))
	}
	return datareq
}
func ListrequserToRes(data []request.RequestUser) (datareq []respon.ResponseUser) {
	for _, val := range data {
		datareq = append(datareq, ReqtoRepon(val))
	}
	return datareq
}
func ReqtoReponbank(data request.RequestBank) respon.ResponBank {
	return respon.ResponBank{
		NoRekening: data.NoRekening,
		NamaBank:   data.NamaBank,
		IDPlayer:   data.IDPlayer,
	}
}
func RequesbankToModel(data request.RequestBank) model.Bank {
	return model.Bank{

		NoRekening: data.NoRekening,
		NamaBank:   data.NamaBank,
		IDPlayer:   data.IDPlayer,
	}
}
func ModelbanktoReq(data *model.Bank) request.RequestBank {
	return request.RequestBank{
		Id:         int(data.ID),
		NoRekening: data.NoRekening,
		NamaBank:   data.NamaBank,
		IDPlayer:   data.IDPlayer,
	}
}
func ReqtoReponwallet(data request.RequestWallet) respon.ResponWallet {
	return respon.ResponWallet{
		Id:          data.Id,
		Nama_dompet: data.Nama_dompet,
		Saldo:       data.Saldo,
		IDPlayer:    data.IDPlayer,
	}
}
func RequeswalletToModel(data request.RequestWallet) model.Wallet {
	return model.Wallet{

		NamaDompet: data.Nama_dompet,
		Saldo:      data.Saldo,
		IDPlayer:   data.IDPlayer,
	}
}
func ModelwallettoReq(data *model.Wallet) request.RequestWallet {
	return request.RequestWallet{
		Id:          int(data.ID),
		Nama_dompet: data.NamaDompet,
		Saldo:       data.Saldo,
		IDPlayer:    data.IDPlayer,
	}
}
