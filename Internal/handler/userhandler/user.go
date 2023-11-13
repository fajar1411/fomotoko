package userhandler

import (
	"context"
	"encoding/json"
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"
	"test/middlewares"
	"test/redist"

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

func (hu *HandlerUser) RegisterUser(e echo.Context) error {
	requestRegister := request.RequestUser{}

	binderr := e.Bind(&requestRegister)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)

	_, errexist := client.HExists(context.Background(), "users", requestRegister.Email).Result()
	if errexist != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errexist.Error(), http.StatusInternalServerError, true))
	}

	data, errservice := hu.um.RegisterUser(requestRegister)
	userDataJSON, err := json.Marshal(data)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(err.Error(), http.StatusInternalServerError, true))
	}
	errredist := client.HSet(context.Background(), "users", data.Email, userDataJSON).Err()
	if errredist != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errredist.Error(), http.StatusInternalServerError, true))
	}
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoRepon(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}

// AllUser implements handlecontract.HandleUser.
func (hc *HandlerUser) AllUser(e echo.Context) error {
	dataservice, errservice := hc.um.AllUser()
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ListrequserToRes(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))

}

func (hu *HandlerUser) Player(e echo.Context) error {
	id := middlewares.ExtractTokenId(e)

	if id <= 0 {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("id tidak ada", http.StatusUnauthorized, true))

	}
	dataservice, errservice := hu.um.Player(id)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, http.StatusInternalServerError, true))
	}
	respon := query.ReqProfiletores(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, true))
}

// Filter implements handlecontract.HandleUser.
func (hu *HandlerUser) Filter(e echo.Context) error {
	nama := e.QueryParam("nama")
	norek := e.QueryParam("norek")

	dataservice, errservice := hu.um.Filter(nama, norek)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, http.StatusInternalServerError, true))
	}
	respon := query.FilterequserToRes(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, true))

}
