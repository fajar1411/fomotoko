package userservice

import (
	"errors"
	"test/bycripts"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"
	"test/validasi"
	"time"

	"github.com/go-playground/validator"
)

type ServicesCase struct {
	ru       repocontract.RepoUser
	validate *validator.Validate
}

func NewServiceUser(ru repocontract.RepoUser) servicecontract.ServiceCase {
	return &ServicesCase{
		ru:       ru,
		validate: validator.New(),
	}
}

func (sc *ServicesCase) RegisterAdmin(newRequest request.RequestUser) (data request.RequestUser, err error) {
	validerr := sc.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestUser{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	newRequest.CreatedAt = time.Now()
	haspw := bycripts.Bcript(newRequest.Password)
	newRequest.Password = haspw

	newRequest.Role = "admin"

	datarepo, errrepo := sc.ru.RegisterAdmin(newRequest)

	if errrepo != nil {
		return request.RequestUser{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

func (sc *ServicesCase) RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error) {
	validerr := sc.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestUser{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	newRequest.CreatedAt = time.Now()
	haspw := bycripts.Bcript(newRequest.Password)
	newRequest.Password = haspw

	datarepo, errrepo := sc.ru.RegisterUser(newRequest)

	if errrepo != nil {
		return request.RequestUser{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

// // Register implements servicecontract.ServiceCase.
// func (sc *ServicesCase) Register(newRequest request.RequestUser) (data request.RequestUser, err error) {

// 	validerr := sc.validate.Struct(newRequest)
// 	if validerr != nil {

// 		return request.RequestUser{}, errors.New(validasi.ValidationErrorHandle(validerr))
// 	}
// 	newRequest.CreatedAt = time.Now()
// 	haspw := bycripts.Bcript(newRequest.Password)
// 	newRequest.Password = haspw

// 	datarepo, errrepo := sc.ru.Register(newRequest)

// 	if errrepo != nil {
// 		return request.RequestUser{}, errors.New(errrepo.Error())
// 	}

// 	return datarepo, nil
// }
