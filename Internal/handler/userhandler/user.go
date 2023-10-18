package userhandler

import (
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerUser struct {
	um servicecontract.ServiceCase
}

func NewHandleUser(um servicecontract.ServiceCase) handlecontract.HandleUser {
	return &HandlerUser{
		um: um,
	}
}

func (hu *HandlerUser) RegisterAdmin(e echo.Context) error {
	requestRegister := request.RequestUser{}

	binderr := e.Bind(&requestRegister)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	data, errservice := hu.um.RegisterAdmin(requestRegister)
	// fmt.Print("ini data handler", data)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoRepon(data)
	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}

func (hu *HandlerUser) RegisterUser(e echo.Context) error {
	requestRegister := request.RequestUser{}

	binderr := e.Bind(&requestRegister)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	data, errservice := hu.um.RegisterUser(requestRegister)
	// fmt.Print("ini data handler", data)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoRepon(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}
