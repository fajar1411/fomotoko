package servicecontract

import "test/domain/request"

type ServiceCase interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	RegisterAdmin(newRequest request.RequestUser) (data request.RequestUser, err error)
}

type ServiceLogin interface {
	LoginAdmin(email string, password string) (string, request.RequestUser, error)
	LoginUser(email string, password string) (string, request.RequestUser, error)
}
type ServiceBarang interface {
	AddBarang(newreq request.RequestBarang) (data request.RequestBarang, err error)
}
type ServiceOrder interface {
	AddOrder(Id int, newreq request.RequestOrder) (data request.RequestOrder, err error)
}
