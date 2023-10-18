package barangservice

import (
	"errors"
	"strconv"
	"test/domain/contract/repocontract"
	"test/domain/contract/servicecontract"
	"test/domain/request"

	"github.com/go-playground/validator"
)

type ServicesOrder struct {
	ro       repocontract.RepoOrder
	rb       repocontract.RepoBarang
	validate *validator.Validate
}

func NewServiceOrder(rb repocontract.RepoBarang, ro repocontract.RepoOrder) servicecontract.ServiceOrder {
	return &ServicesOrder{
		rb:       rb,
		ro:       ro,
		validate: validator.New(),
	}
}

func (so *ServicesOrder) AddOrder(userid int, Id int, newreq request.RequestOrder) (data request.RequestOrder, err error) {
	newreq.UserId = userid
	barang, erridbarang := so.rb.IdBarang(Id)

	if erridbarang != nil {
		return request.RequestOrder{}, erridbarang
	}
	newreq.Status = "terjual"
	newreq.BarangId = barang.Id
	newreq.NamaBarang = barang.NamaBarang
	newreq.Harga = barang.Harga
	cnvharga, errcnv := strconv.Atoi(barang.Harga)

	if newreq.Quantity > barang.Stok {
		return request.RequestOrder{}, errors.New("Stok Tidak Cukup Untuk Di beli")
	}

	if errcnv != nil {
		return request.RequestOrder{}, errcnv
	}
	totalbarang := cnvharga * newreq.Quantity

	cnvtotal := strconv.Itoa(totalbarang)

	newreq.TotalOrder = cnvtotal

	datarepo, errrepo := so.ro.AddOrder(newreq)

	if errrepo != nil {
		return request.RequestOrder{}, errrepo
	}
	return datarepo, nil

}
