package repobank

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"

	"gorm.io/gorm"
)

type Repobank struct {
	db *gorm.DB
}

func NewRepoBank(db *gorm.DB) repocontract.RepoBank {
	return &Repobank{
		db: db,
	}
}

func (rb *Repobank) CreateBank(newRequest request.RequestBank) (data request.RequestBank, err error) {
	datareqtomodelbank := query.RequesbankToModel(newRequest)

	tx := rb.db.Create(&datareqtomodelbank)

	if tx.Error != nil {
		return request.RequestBank{}, tx.Error
	}
	datamodeltoreq := query.ModelbanktoReq(&datareqtomodelbank)
	return datamodeltoreq, nil
}

func (rb *Repobank) IduserExist(iduser int) (data request.RequestBank, err error) {
	var activ model.Bank

	tx := rb.db.Raw("Select banks.id, banks.no_rekening, banks.nama_bank, bankid_player from banks WHERE banks.id_player= ? ", iduser).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestBank{}, tx.Error
	}
	var activcore = query.ModelbanktoReq(&activ)
	return activcore, nil
}

// norekExist implements repocontract.RepoBank.
func (rb *Repobank) NorekExist(no string) (data request.RequestBank, err error) {
	var activ model.Bank

	tx := rb.db.Raw("Select banks.id, banks.no_rekening, banks.nama_bank, bankid_player from banks WHERE banks.no_rekening= ? ", no).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestBank{}, tx.Error
	}
	var activcore = query.ModelbanktoReq(&activ)
	return activcore, nil
}
