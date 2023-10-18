package repocontract

import "test/domain/request"

type RepoUser interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	RegisterAdmin(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllAdmin() (data []request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	EmailadminExist(email string) (data request.RequestUser, err error)
	EmaiuserExist(email string) (data request.RequestUser, err error)
}
type RepoLogin interface {
	LoginAdmin(email string, password string) (string, request.RequestUser, error)
	LoginUser(email string, password string) (string, request.RequestUser, error)
}
type RepoBarang interface {
	AddBarang(newreq request.RequestBarang) (data request.RequestBarang, err error)
	IdBarang(id int) (data request.RequestBarang, err error)
}
type RepoOrder interface {
	AddOrder(newreq request.RequestOrder) (data request.RequestOrder, err error)
}
