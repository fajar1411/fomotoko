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

func ReqtoRepon(data request.RequestUser) respon.Responuserandwallet {
	return respon.Responuserandwallet{
		Id:            data.Id,
		Password:      data.Password,
		Email:         data.Email,
		Name:          data.Name,
		Nama_dompet:   data.Nama_dompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
	}
}
func ReqtoRespon(data request.RequestUser) respon.ResponseUser {
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
		datareq = append(datareq, ReqtoRespon(val))
	}
	return datareq
}
func ReqtoReponbank(data request.RequestBank) respon.ResponBank {
	return respon.ResponBank{
		NoRekening:   data.NoRekening,
		NamaBank:     data.NamaBank,
		IDPlayer:     data.IDPlayer,
		NamaRekening: data.NamaRekening,
	}
}
func RequesbankToModel(data request.RequestBank) model.Bank {
	return model.Bank{
		NamaRekening: data.NamaRekening,
		NoRekening:   data.NoRekening,
		NamaBank:     data.NamaBank,
		IDPlayer:     data.IDPlayer,
	}
}
func ModelbanktoReq(data *model.Bank) request.RequestBank {
	return request.RequestBank{
		Id:           int(data.ID),
		NamaRekening: data.NamaRekening,
		NoRekening:   data.NoRekening,
		NamaBank:     data.NamaBank,
		IDPlayer:     data.IDPlayer,
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
func ReqtoRepontopup(data request.RequestTopUp) respon.ResponTopUp {
	return respon.ResponTopUp{
		Id:            data.Id,
		IDPlayer:      data.IDPlayer,
		Saldo:         data.Saldo,
		OrderId:       data.OrderId,
		PaymentMethod: data.PaymentMethod,
		Status:        data.Status,
		IDWallet:      data.IDWallet,
	}
}
func Reqtomodeltopup(data request.RequestTopUp) model.TopUp {
	return model.TopUp{

		Saldo:         data.Saldo,
		OrderId:       data.OrderId,
		PaymentMethod: data.PaymentMethod,
		Status:        data.Status,
		IDWallet:      data.IDWallet,
		IdUser:        data.IDPlayer,
	}
}
func ModeltoReqtopup(data *model.TopUp) request.RequestTopUp {
	return request.RequestTopUp{
		Id:            int(data.ID),
		Saldo:         data.Saldo,
		OrderId:       data.OrderId,
		PaymentMethod: data.PaymentMethod,
		Status:        data.Status,
		IDWallet:      data.IDWallet,
		IDPlayer:      data.IdUser,
	}
}

func Requsertomodelwallet(data request.RequestUser) model.Wallet {
	return model.Wallet{

		NamaDompet:    data.Nama_dompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
		IDPlayer:      uint(data.Id),
	}
}
func Modelwallettorequser(data model.Wallet) request.RequestUser {
	return request.RequestUser{

		Nama_dompet:   data.NamaDompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
		Id:            int(data.IDPlayer),
	}
}
func Reqwalletanduser(data request.RequestUserDanwallet) request.RequestUser {
	return request.RequestUser{
		Id:            data.Id,
		Password:      data.Password,
		Email:         data.Email,
		Name:          data.Email,
		Nama_dompet:   data.Nama_dompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
	}
}
func ModelProfiletoReqprofile(data model.Profile) request.ReqProfile {
	return request.ReqProfile{
		IdUser:        data.ID,
		Email:         data.Email,
		Nama:          data.Nama,
		NamaDompet:    data.NamaDompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
		NoRekening:    data.NoRekening,
	}
}
func ReqProfiletores(data request.ReqProfile) respon.ResProfile {
	return respon.ResProfile{
		IdUser:        data.IdUser,
		Email:         data.Email,
		Nama:          data.Nama,
		NamaDompet:    data.NamaDompet,
		AccountWallet: data.AccountWallet,
		Saldo:         data.Saldo,
		NoRekening:    data.NoRekening,
	}
}
func FilterequserToRes(data []request.ReqProfile) (datareq []respon.ResProfile) {
	for _, val := range data {
		datareq = append(datareq, ReqProfiletores(val))
	}
	return datareq
}
func FiltemodeluserToReq(data []model.Profile) (datareq []request.ReqProfile) {
	for _, val := range data {
		datareq = append(datareq, ModelProfiletoReqprofile(val))
	}
	return datareq
}
