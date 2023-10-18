package repouser

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"

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

func (ru *RepoUser) RegisterAdmin(newRequest request.RequestUser) (data request.RequestUser, err error) {
	getalladmin, errall := ru.AllAdmin()

	if errall != nil {
		return request.RequestUser{}, errall
	}

	lenadmin := len(getalladmin)

	datareqtomodeladmin := query.RequadminToModel(newRequest)

	if lenadmin <= 0 || lenadmin > 0 {
		lenadmin += 1

		datareqtomodeladmin.Id = lenadmin

	}
	_, errexist := ru.EmailadminExist(datareqtomodeladmin.Email)

	if errexist == nil {
		return request.RequestUser{}, errors.New("Email Sudah Terdaftar")
	}
	tx := ru.db.Create(&datareqtomodeladmin)

	if tx.Error != nil {
		return request.RequestUser{}, tx.Error
	}
	datamodeltoreq := query.ModeltoReq(datareqtomodeladmin)
	return datamodeltoreq, nil

}

func (ru *RepoUser) RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error) {
	getalluser, errall := ru.AllUser()

	if errall != nil {
		return request.RequestUser{}, errall
	}

	lenuser := len(getalluser)

	datareqtomodeluser := query.RequuserToModel(newRequest)

	if lenuser <= 0 || lenuser > 0 {
		lenuser += 1

		datareqtomodeluser.Id = lenuser

	}
	_, errexist := ru.EmaiuserExist(datareqtomodeluser.Email)

	if errexist == nil {
		return request.RequestUser{}, errors.New("Email Sudah Terdaftar")
	}
	tx := ru.db.Create(&datareqtomodeluser)

	if tx.Error != nil {
		return request.RequestUser{}, tx.Error
	}
	datamodeltoreq := query.ModelusertoReq(datareqtomodeluser)
	return datamodeltoreq, nil
}

func (ru *RepoUser) AllAdmin() (data []request.RequestUser, err error) {
	var activ []model.Admin
	tx := ru.db.Raw("Select admins.id, admins.password, admins.email,admins.nama, admins.roles from admins").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelToReq(activ)
	return dtmdlttoreq, nil
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

func (ru *RepoUser) EmailadminExist(email string) (data request.RequestUser, err error) {
	var activ model.Admin

	tx := ru.db.Raw("Select admins.id, admins.password, admins.email, admins.nama,admins.roles from admins WHERE admins.email= ? ", email).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(activ)
	return activcore, nil
}

func (ru *RepoUser) EmaiuserExist(email string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.password, users.email, users.nama from users WHERE users.email= ? ", email).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModelusertoReq(activ)
	return activcore, nil
}
