package orderhandler

import (
	"net/http"
	"strconv"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"
	middlewares "test/middleware"

	echo "github.com/labstack/echo/v4"
)

type HandlerOrder struct {
	sb servicecontract.ServiceOrder
}

func NewHandlOrder(sb servicecontract.ServiceOrder) handlecontract.HandleOrder {
	return &HandlerOrder{
		sb: sb,
	}
}

func (ho *HandlerOrder) AddOrder(e echo.Context) error {
	reqorder := request.RequestOrder{}
	userid := middlewares.ExtractTokenTeamId(e)
	idbarang := e.Param("idbarang")
	cnv, errcnv := strconv.Atoi(idbarang)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv.Error(), http.StatusBadRequest, true))
	}
	errbind := e.Bind(&reqorder)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := ho.sb.AddOrder(userid, cnv, reqorder)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	responorder := query.ReqtoResponOrder(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(responorder, http.StatusCreated, false))
}
