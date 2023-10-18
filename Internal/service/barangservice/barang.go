package barangservice

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"
	"test/validasi"
	"time"

	"github.com/go-playground/validator"
)

type ServicesBarang struct {
	rb       repocontract.RepoBarang
	validate *validator.Validate
}

func NewServiceBarang(rb repocontract.RepoBarang) servicecontract.ServiceBarang {
	return &ServicesBarang{
		rb:       rb,
		validate: validator.New(),
	}
}

func (sb *ServicesBarang) AddBarang(newreq request.RequestBarang) (data request.RequestBarang, err error) {

	validerr := sb.validate.Struct(newreq)
	if validerr != nil {

		return request.RequestBarang{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	if newreq.Stok <= 0 {
		return request.RequestBarang{}, errors.New("Input Stok Harus diatas 0 tidak Boleh Minus maupun 0")
	}
	newreq.CreatedAt = time.Now()

	datarepo, errrepo := sb.rb.AddBarang(newreq)

	if errrepo != nil {
		return request.RequestBarang{}, errrepo
	}
	return datarepo, nil
}
