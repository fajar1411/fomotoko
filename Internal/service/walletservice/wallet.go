package walletservice

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"
	"test/validasi"

	"github.com/go-playground/validator"
)

type Serviceswallet struct {
	ru       repocontract.RepoUser
	rw       repocontract.RepoWallet
	validate *validator.Validate
}

func NewServiceWallet(ru repocontract.RepoUser, rw repocontract.RepoWallet) servicecontract.ServiceWallet {
	return &Serviceswallet{
		ru:       ru,
		rw:       rw,
		validate: validator.New(),
	}
}

func (sw *Serviceswallet) CreateWallet(newRequest request.RequestWallet, email string) (data request.RequestWallet, err error) {
	validerr := sw.validate.Struct(newRequest)
	if validerr != nil {

		return data, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	emailexist, errexist := sw.ru.EmaiuserExist(email)

	if errexist != nil {
		return data, errors.New(errexist.Error())
	}
	newRequest.IDPlayer = uint(emailexist.Id)
	newRequest.Saldo = 0
	_, errdompetuser := sw.rw.IduserExist(int(newRequest.IDPlayer))
	if errdompetuser == nil {
		return data, errors.New("Anda sudah Membuat dompet")
	}
	datarepo, errrepo := sw.rw.CreateWallet(newRequest)

	if errrepo != nil {
		return data, errors.New(errrepo.Error())
	}

	return datarepo, nil
}
