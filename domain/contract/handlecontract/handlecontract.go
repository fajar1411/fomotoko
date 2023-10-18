package handlecontract

import "github.com/labstack/echo/v4"

type HandleUser interface {
	RegisterAdmin(e echo.Context) error
	RegisterUser(e echo.Context) error
}

type HandleLogin interface {
	LoginAdmin(e echo.Context) error
	LoginUser(e echo.Context) error
}

type HandleBarang interface {
	AddBarang(e echo.Context) error
}

type HandleOrder interface {
	AddOrder(e echo.Context) error
}
