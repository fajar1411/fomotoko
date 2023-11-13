package loginhandler

import (
	"context"
	"net/http"
	"test/domain/contract/handlecontract"
	"test/domain/contract/servicecontract"
	"test/domain/query"
	"test/domain/request"
	"test/helper"
	"test/redist"

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

func (hl *HandlerLogin) LoginUser(e echo.Context) error {
	reques := request.RequestUser{}

	errbind := e.Bind(&reques)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}
	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)
	token, dataservice, errservice := hl.sl.LoginUser(reques.Email, reques.Password)
	err := client.Set(context.Background(), "login", token, redist.Expiration).Err()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(err.Error(), http.StatusInternalServerError, true))
	}
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	e.Response().Header().Set("Authorization", token)
	respon := query.ReqtoResponLogin(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hl *HandlerLogin) LogoutUser(e echo.Context) error {
	token := e.Request().Header.Get("Authorization")

	if token == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Token not provided", http.StatusUnauthorized, true))
	}

	client := redist.RedisClient()
	defer redist.CloseRedisConnection(client)

	// Hapus token dari Redis
	err := client.Del(context.Background(), "login", token).Err()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(err.Error(), http.StatusInternalServerError, true))
	}

	return e.JSON(http.StatusOK, helper.GetResponse("Logout successful", http.StatusOK, false))
}
