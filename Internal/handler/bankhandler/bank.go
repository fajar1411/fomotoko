package bankhandler

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

type HandlerBank struct {
	sb servicecontract.ServiceBank
}

func NewHandleBank(sb servicecontract.ServiceBank) handlecontract.HandleBank {
	return &HandlerBank{
		sb: sb,
	}
}

func (hb *HandlerBank) CreateBank(e echo.Context) error {
	RequestBank := request.RequestBank{}
	binderr := e.Bind(&RequestBank)
	email := middlewares.ExtractTokenEmail(e)
	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	service, errservice := hb.sb.CreateBank(RequestBank, email)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoReponbank(service)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}
