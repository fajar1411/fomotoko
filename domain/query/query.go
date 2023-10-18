package query

import (
	"test/domain/model"
	"test/domain/request"
	"test/domain/respon"
)

func RequadminToModel(data request.RequestUser) model.Admin {
	return model.Admin{
		Id:        data.Id,
		Roles:     data.Role,
		Email:     data.Email,
		Nama:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}
func RequuserToModel(data request.RequestUser) model.User {
	return model.User{
		Id:        data.Id,
		Email:     data.Email,
		Nama:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}
func ModelusertoReq(data model.User) request.RequestUser {
	return request.RequestUser{
		Id:        data.Id,
		Email:     data.Email,
		Name:      data.Nama,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}

func ModeltoReq(data model.Admin) request.RequestUser {
	return request.RequestUser{
		Id:        data.Id,
		Role:      data.Roles,
		Email:     data.Email,
		Name:      data.Nama,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}

func ReqtoRepon(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Id:        data.Id,
		Role:      data.Role,
		Email:     data.Email,
		Name:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}

func ReqtoResponLogin(data request.RequestUser, token string) respon.ResponseUser {
	return respon.ResponseUser{
		Id:        data.Id,
		Role:      data.Role,
		Email:     data.Email,
		Name:      data.Name,
		Token:     token,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}
func ListModelToReq(data []model.Admin) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModeltoReq(val))
	}
	return datareq
}
func ListModeluserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModelusertoReq(val))
	}
	return datareq
}

func ReqtoResponBarang(data request.RequestBarang) respon.ResponseBarang {
	return respon.ResponseBarang{

		NamaBarang: data.NamaBarang,
		Stok:       data.Stok,
		Harga:      data.Harga,
	}
}

func ReqtoModelBarang(data request.RequestBarang) model.Barang {
	return model.Barang{

		NamaBarang: data.NamaBarang,
		Stok:       data.Stok,
		Harga:      data.Harga,
	}
}
func ModeltoReqBarang(data model.Barang) request.RequestBarang {
	return request.RequestBarang{

		NamaBarang: data.NamaBarang,
		Stok:       data.Stok,
		Harga:      data.Harga,
	}
}
