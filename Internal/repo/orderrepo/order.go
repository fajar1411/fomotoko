package orderrepo

import (
	"test/domain/contract/repocontract"
	"test/domain/query"
	"test/domain/request"

	"gorm.io/gorm"
)

type RepoOrder struct {
	db *gorm.DB
}

func NewRepoOrder(db *gorm.DB) repocontract.RepoOrder {
	return &RepoOrder{
		db: db,
	}
}

func (ro *RepoOrder) AddOrder(newreq request.RequestOrder) (data request.RequestOrder, err error) {
	reqtomodel := query.ReqtoModelorder(newreq)
	tx := ro.db.Create(&reqtomodel)

	if tx.Error != nil {
		return request.RequestOrder{}, tx.Error
	}
	modeltoreq := query.Modeltoreqorder(reqtomodel)

	return modeltoreq, nil
}
