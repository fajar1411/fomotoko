package faktoryandroute

import (
	uh "test/internal/handler/userhandler"
	ru "test/internal/repo/repouser"
	us "test/internal/service/userservice"
	"test/middlewares"

	lh "test/internal/handler/loginhandler"
	rl "test/internal/repo/repologin"
	ls "test/internal/service/loginservice"

	bh "test/internal/handler/bankhandler"
	rb "test/internal/repo/repobank"
	bs "test/internal/service/bankservice"

	wh "test/internal/handler/wallethandler"
	rw "test/internal/repo/repowallet"
	ws "test/internal/service/walletservice"

	th "test/internal/handler/topuphandler"
	rt "test/internal/repo/repotopup"
	ts "test/internal/service/topupservice"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpm := ru.NewRepoUser(db)
	ucmhsw := us.NewServiceUser(rpm)
	hndlmhs := uh.NewHandleUser(ucmhsw)
	usergrup := e.Group("/player")
	usergrup.POST("/register", hndlmhs.RegisterUser)
	usergrup.GET("", hndlmhs.AllUser)
	usergrup.GET("/profile", hndlmhs.Player, middlewares.JWTMiddleware())
	usergrup.GET("/filterlist", hndlmhs.Filter)

	rpl := rl.NewRepoLogin(db)
	servicelogin := ls.NewServiceLogin(rpl, rpm)
	handlelogin := lh.NewHandlLogin(servicelogin)
	logingrup := e.Group("/player")

	logingrup.POST("/login", handlelogin.LoginUser)
	logingrup.POST("/logout", handlelogin.LogoutUser)

	rpb := rb.NewRepoBank(db)
	servicebank := bs.NewServiceBank(rpm, rpb)
	handlebank := bh.NewHandleBank(servicebank)
	bankgrup := e.Group("/player")

	bankgrup.POST("/addbank", handlebank.CreateBank, middlewares.JWTMiddleware())
	rpw := rw.NewRepoWallet(db)
	servicewallet := ws.NewServiceWallet(rpm, rpw)
	handlewallet := wh.NewHandleWallet(servicewallet)
	walletgrup := e.Group("/player")

	walletgrup.POST("/addwallet", handlewallet.CreateWallet, middlewares.JWTMiddleware())
	rpt := rt.NewRepoTopup(db)
	servictopup := ts.NewServiceTopUp(rpt, rpm, rpw)
	handletopup := th.NewHandleTopUp(servictopup)
	topupgrup := e.Group("/player")

	topupgrup.POST("/addtopup", handletopup.CreateTopUp, middlewares.JWTMiddleware())

}
