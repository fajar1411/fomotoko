package loginhandler

import (
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerLogin struct {
	sl servicecontract.ServiceLogin
}

func NewHandlLogin(sl servicecontract.ServiceLogin) handlecontract.HandleLogin {
	return &HandlerLogin{
		sl: sl,
	}
}

// LoginAdmin implements handlecontract.HandleLogin.
func (hl *HandlerLogin) LoginAdmin(e echo.Context) error {
	reques := request.RequestUser{}

	errbind := e.Bind(&reques)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}

	token, dataservice, errservice := hl.sl.LoginAdmin(reques.Email, reques.Password)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqtoResponLogin(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))

}

// LoginUser implements handlecontract.HandleLogin.
func (hl *HandlerLogin) LoginUser(e echo.Context) error {
	reques := request.RequestUser{}

	errbind := e.Bind(&reques)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}

	token, dataservice, errservice := hl.sl.LoginUser(reques.Email, reques.Password)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqtoResponLogin(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}
