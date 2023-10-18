package barangrepo

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"

	"gorm.io/gorm"
)

type RepoBarang struct {
	db *gorm.DB
}

func NewRepoBarang(db *gorm.DB) repocontract.RepoBarang {
	return &RepoBarang{
		db: db,
	}
}

func (rb *RepoBarang) AddBarang(newreq request.RequestBarang) (data request.RequestBarang, err error) {
	reqtomodel := query.ReqtoModelBarang(newreq)

	tx := rb.db.Create(&reqtomodel)

	if tx.Error != nil {
		return request.RequestBarang{}, tx.Error
	}
	modeltoreq := query.ModeltoReqBarang(reqtomodel)
	return modeltoreq, nil
}

func (rb *RepoBarang) IdBarang(id int) (data request.RequestBarang, err error) {
	var activ model.Barang

	tx := rb.db.Raw("Select barangs.id, barangs.nama_barang, barangs.stok, barangs.harga from barangs WHERE barangs.id= ? ", id).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestBarang{}, tx.Error
	}
	var activcore = query.ModeltoReqBarang(activ)
	return activcore, nil
}
