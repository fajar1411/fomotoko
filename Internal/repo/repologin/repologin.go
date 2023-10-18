package repologin

import (
	"errors"
	"test/domain/contract/repocontract"
	"test/domain/model"
	"test/domain/query"
	"test/domain/request"
	middlewares "test/middleware"

	"gorm.io/gorm"
)

type Repologin struct {
	db *gorm.DB
}

func (*Repologin) Login(email string, password string) (string, request.RequestUser, error) {
	panic("unimplemented")
}

func NewRepoLogin(db *gorm.DB) repocontract.RepoLogin {
	return &Repologin{
		db: db,
	}
}

func (rl *Repologin) LoginAdmin(email string, password string) (string, request.RequestUser, error) {
	userdata := model.Admin{}

	tx := rl.db.Where("email = ?", email).First(&userdata)
	if tx.Error != nil {
		return "", request.RequestUser{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateTokenTeam(userdata.Id, userdata.Roles)

	if errtoken != nil {
		return "", request.RequestUser{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModeltoReq(userdata)
	return createtoken, datamodeltoreq, nil
}

func (rl *Repologin) LoginUser(email string, password string) (string, request.RequestUser, error) {
	userdata := model.User{}

	tx := rl.db.Where("email = ?", email).First(&userdata)
	if tx.Error != nil {
		return "", request.RequestUser{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateTokenTeam(userdata.Id, "")

	if errtoken != nil {
		return "", request.RequestUser{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModelusertoReq(userdata)
	return createtoken, datamodeltoreq, nil
}
