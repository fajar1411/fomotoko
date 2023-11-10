package wallethandler

import (
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"
	"test/middlewares"

	echo "github.com/labstack/echo/v4"
)

type Handlerwallet struct {
	sw servicecontract.ServiceWallet
}

func NewHandleWallet(sw servicecontract.ServiceWallet) handlecontract.HandleWallet {
	return &Handlerwallet{
		sw: sw,
	}
}

// CreateWallet implements handlecontract.HandleWallet.
func (hw *Handlerwallet) CreateWallet(e echo.Context) error {
	Requestwallet := request.RequestWallet{}
	binderr := e.Bind(&Requestwallet)
	email := middlewares.ExtractTokenEmail(e)
	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	service, errservice := hw.sw.CreateWallet(Requestwallet, email)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoReponwallet(service)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}
