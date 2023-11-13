package repotopup

import (
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"

	"gorm.io/gorm"
)

type Repotopup struct {
	db *gorm.DB
}

func NewRepoTopup(db *gorm.DB) repocontract.RepoTopUp {
	return &Repotopup{
		db: db,
	}
}

// CreateTopUp implements repocontract.RepoTopUp.
func (rt *Repotopup) CreateTopUp(newRequest request.RequestTopUp) (data request.RequestTopUp, err error) {
	datareqtomodel := query.Reqtomodeltopup(newRequest)

	tx := rt.db.Create(&datareqtomodel)
	wallet := model.Wallet{}

	if tx.Error != nil {
		return data, tx.Error
	}

	if newRequest.Status == "sukses payment" {
		txupdated := rt.db.Model(&wallet).Where("id =?", datareqtomodel.IDWallet).Update("saldo", newRequest.UpdateSaldo)
		if txupdated.Error != nil {
			return request.RequestTopUp{}, txupdated.Error
		}
	}

	datamodeltoreq := query.ModeltoReqtopup(&datareqtomodel)

	return datamodeltoreq, nil
}

// CreatePayment implements repocontract.RepoTopUp.
func (*Repotopup) CreatePayment(newRequest request.RequestTopUp) (data request.RequestTopUp, err error) {
	panic("unimplemented")
}
