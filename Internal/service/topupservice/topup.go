package topupervice

import (
	"errors"
	"fmt"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"
	"test/helper"
	"test/validasi"

	"github.com/go-playground/validator"
)

type Servicestopup struct {
	ru       repocontract.RepoUser
	rw       repocontract.RepoWallet
	rt       repocontract.RepoTopUp
	validate *validator.Validate
}

func NewServiceTopUp(rt repocontract.RepoTopUp, ru repocontract.RepoUser, rw repocontract.RepoWallet) servicecontract.ServiceTopUp {
	return &Servicestopup{
		rt:       rt,
		ru:       ru,
		rw:       rw,
		validate: validator.New(),
	}
}

func (st *Servicestopup) CreateTopUp(newRequest request.RequestTopUp, email string) (data request.RequestTopUp, err error) {
	validerr := st.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestTopUp{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	emailexist, errexist := st.ru.EmaiuserExist(email)

	if errexist != nil {
		return request.RequestTopUp{}, errors.New(errexist.Error())
	}
	walletexist, errwallet := st.rw.IduserExist(emailexist.Id)
	if errwallet != nil {
		return request.RequestTopUp{}, errors.New(errwallet.Error())
	}
	newRequest.IDPlayer = uint(emailexist.Id)
	newRequest.IDWallet = uint(walletexist.Id)
	newRequest.Status = "sukses payment"
	newRequest.UpdateSaldo = walletexist.Saldo + newRequest.Saldo
	randString := helper.FileName(5)
	newRequest.OrderId = "order-" + randString
	datarepo, errrepo := st.rt.CreateTopUp(newRequest)

	fmt.Print("ini data service bro", datarepo)

	if errrepo != nil {
		return request.RequestTopUp{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

// CreatePayment implements servicecontract.ServiceTopUp.
func (st *Servicestopup) CreatePayment(newRequest request.RequestTopUp) (data request.RequestTopUp, err error) {
	panic("unimplemented")
}
