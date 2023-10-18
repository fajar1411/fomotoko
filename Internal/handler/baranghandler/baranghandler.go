package baranghandler

import (
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"
	middlewares "test/middleware"

	echo "github.com/labstack/echo/v4"
)

type HandlerBarang struct {
	sb servicecontract.ServiceBarang
}

func NewHandlBarang(sb servicecontract.ServiceBarang) handlecontract.HandleBarang {
	return &HandlerBarang{
		sb: sb,
	}
}

// AddBarang implements handlecontract.HandleBarang.
func (hb *HandlerBarang) AddBarang(e echo.Context) error {
	reqbarang := request.RequestBarang{}

	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	errbind := e.Bind(&reqbarang)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}

	data, errservice := hb.sb.AddBarang(reqbarang)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	responBarang := query.ReqtoResponBarang(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(responBarang, http.StatusCreated, false))
}
