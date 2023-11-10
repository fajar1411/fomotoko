package bankservice

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"
	"test/validasi"

	"github.com/go-playground/validator"
)

type Servicesbank struct {
	ru       repocontract.RepoUser
	rb       repocontract.RepoBank
	validate *validator.Validate
}

func NewServiceBank(ru repocontract.RepoUser, rb repocontract.RepoBank) servicecontract.ServiceBank {
	return &Servicesbank{
		ru:       ru,
		rb:       rb,
		validate: validator.New(),
	}
}

func (sb *Servicesbank) CreateBank(newRequest request.RequestBank, email string) (data request.RequestBank, err error) {
	validerr := sb.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestBank{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	emailexist, errexist := sb.ru.EmaiuserExist(email)

	if errexist != nil {
		return request.RequestBank{}, errors.New(errexist.Error())
	}
	newRequest.IDPlayer = uint(emailexist.Id)
	_, errrek := sb.rb.NorekExist(newRequest.NoRekening)
	if errrek == nil {
		return request.RequestBank{}, errors.New("no rekening sudah ada")
	}
	datarepo, errrepo := sb.rb.CreateBank(newRequest)

	if errrepo != nil {
		return request.RequestBank{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}
