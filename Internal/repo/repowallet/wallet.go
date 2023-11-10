package repowallet

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"

	"gorm.io/gorm"
)

type Repowallet struct {
	db *gorm.DB
}

func NewRepoWallet(db *gorm.DB) repocontract.RepoWallet {
	return &Repowallet{
		db: db,
	}
}

func (rw *Repowallet) CreateWallet(newRequest request.RequestWallet) (data request.RequestWallet, err error) {
	datareqtomodelwallet := query.RequeswalletToModel(newRequest)

	tx := rw.db.Create(&datareqtomodelwallet)

	if tx.Error != nil {
		return data, tx.Error
	}
	datamodeltoreq := query.ModelwallettoReq(&datareqtomodelwallet)
	return datamodeltoreq, nil
}

// IduserExist implements repocontract.RepoWallet.
func (rw *Repowallet) IduserExist(iduser int) (data request.RequestWallet, err error) {
	var activ model.Wallet

	tx := rw.db.Raw("Select wallets.id, wallets.saldo, wallets.id_player, wallets.nama_dompet from wallets WHERE wallets.id_player= ? ", iduser).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return data, tx.Error
	}
	var activcore = query.ModelwallettoReq(&activ)
	return activcore, nil
}
