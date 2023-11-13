package repocontract

import "test/domain/request"

type RepoUser interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	EmaiuserExist(email string) (data request.RequestUser, err error)
	Player(id int) (data request.ReqProfile, err error)
	Filter(nama, norek string) (data []request.ReqProfile, err error)
}
type RepoLogin interface {
	LoginUser(email string, password string) (string, request.RequestUser, error)
}
type RepoBank interface {
	CreateBank(newRequest request.RequestBank) (data request.RequestBank, err error)
	NorekExist(no string) (data request.RequestBank, err error)
	IduserExist(iduser int) (data request.RequestBank, err error)
}
type RepoWallet interface {
	CreateWallet(newRequest request.RequestWallet) (data request.RequestWallet, err error)
	IduserExist(iduser int) (data request.RequestWallet, err error)
}
type RepoTopUp interface {
	CreateTopUp(newRequest request.RequestTopUp) (data request.RequestTopUp, err error)
	CreatePayment(newRequest request.RequestTopUp) (data request.RequestTopUp, err error)
}
