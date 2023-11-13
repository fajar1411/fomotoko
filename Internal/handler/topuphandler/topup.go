package topuphandler

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

type HandlerTopup struct {
	st servicecontract.ServiceTopUp
}

func NewHandleTopUp(st servicecontract.ServiceTopUp) handlecontract.HandleTopUp {
	return &HandlerTopup{
		st: st,
	}
}

func (ht *HandlerTopup) CreateTopUp(e echo.Context) error {
	Requesttopup := request.RequestTopUp{}
	binderr := e.Bind(&Requesttopup)
	email := middlewares.ExtractTokenEmail(e)
	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	service, errservice := ht.st.CreateTopUp(Requesttopup, email)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoRepontopup(service)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}

// CreatePayment implements handlecontract.HandleTopUp.
func (*HandlerTopup) CreatePayment(e echo.Context) error {
	panic("unimplemented")
}
