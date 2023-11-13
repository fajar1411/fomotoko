package repouser

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"
	"test/helper"

	"gorm.io/gorm"
)

type RepoUser struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) repocontract.RepoUser {
	return &RepoUser{
		db: db,
	}
}

func (ru *RepoUser) RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error) {

	datarewaldanuser := request.RequestUserDanwallet{}
	datareqtomodeluser := query.RequuserToModel(newRequest)

	_, errexist := ru.EmaiuserExist(datareqtomodeluser.Email)

	if errexist == nil {
		return request.RequestUser{}, errors.New("Email Sudah Terdaftar")
	}
	tx := ru.db.Create(&datareqtomodeluser)

	if tx.Error != nil {
		return request.RequestUser{}, tx.Error
	}

	datamodeltoreq := query.ModelusertoReq(&datareqtomodeluser)

	datarewaldanuser.Id = datamodeltoreq.Id
	datarewaldanuser.Name = datamodeltoreq.Name
	datarewaldanuser.Email = datamodeltoreq.Email
	datarewaldanuser.Password = datamodeltoreq.Password

	datareqmodelwallet := query.Requsertomodelwallet(datamodeltoreq)
	randString := helper.FileName(5)
	datareqmodelwallet.AccountWallet = datamodeltoreq.Name + randString
	datareqmodelwallet.NamaDompet = datamodeltoreq.Name + "dompet"
	datareqmodelwallet.Saldo = 0
	tx2 := ru.db.Create(&datareqmodelwallet)

	if tx2.Error != nil {
		return request.RequestUser{}, tx.Error
	}
	datamodeltoreqwallet := query.Modelwallettorequser(datareqmodelwallet)

	datarewaldanuser.AccountWallet = datamodeltoreqwallet.AccountWallet
	datarewaldanuser.Nama_dompet = datamodeltoreqwallet.Nama_dompet
	datarewaldanuser.Saldo = datamodeltoreqwallet.Saldo

	datasend := query.Reqwalletanduser(datarewaldanuser)
	return datasend, nil
}

func (ru *RepoUser) AllUser() (data []request.RequestUser, err error) {
	var activ []model.User
	tx := ru.db.Raw("Select users.id, users.password, users.email,users.nama from users").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModeluserToReq(activ)
	return dtmdlttoreq, nil
}

func (ru *RepoUser) EmaiuserExist(email string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.password, users.email, users.nama from users WHERE users.email= ? ", email).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModelusertoReq(&activ)
	return activcore, nil
}
