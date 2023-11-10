package handlecontract

import "github.com/labstack/echo/v4"

type HandleUser interface {
	RegisterUser(e echo.Context) error
	AllUser(e echo.Context) error
}

type HandleLogin interface {
	LoginUser(e echo.Context) error
	LogoutUser(e echo.Context) error
}
type HandleBank interface {
	CreateBank(e echo.Context) error
}
