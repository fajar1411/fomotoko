package servicecontract

import "test/domain/request"

type ServiceCase interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
}

type ServiceLogin interface {
	LoginUser(email string, password string) (string, request.RequestUser, error)
}
type ServiceBank interface {
	CreateBank(newRequest request.RequestBank, email string) (data request.RequestBank, err error)
}
