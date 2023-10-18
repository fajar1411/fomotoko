package faktoryandroute

import (
	uh "test/internal/handler/userhandler"
	ru "test/internal/repo/repouser"
	us "test/internal/service/userservice"

	lh "test/internal/handler/loginhandler"
	rl "test/internal/repo/repologin"
	ls "test/internal/service/loginservice"

	bh "test/internal/handler/baranghandler"
	br "test/internal/repo/barangrepo"
	bs "test/internal/service/barangservice"
	middlewares "test/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpm := ru.NewRepoUser(db)
	ucmhsw := us.NewServiceUser(rpm)
	hndlmhs := uh.NewHandleUser(ucmhsw)
	registergrup := e.Group("/register")
	registergrup.POST("/admin", hndlmhs.RegisterAdmin)
	registergrup.POST("/users", hndlmhs.RegisterUser)

	rpl := rl.NewRepoLogin(db)
	servicelogin := ls.NewServiceLogin(rpl, rpm)
	handlelogin := lh.NewHandlLogin(servicelogin)
	logingrup := e.Group("/login")
	logingrup.POST("/admin", handlelogin.LoginAdmin)
	logingrup.POST("/user", handlelogin.LoginUser)

	rpb := br.NewRepoBarang(db)
	servicebarang := bs.NewServiceBarang(rpb)
	handlebrang := bh.NewHandlBarang(servicebarang)
	barnggrup := e.Group("/barang")
	barnggrup.POST("/addbarang", handlebrang.AddBarang, middlewares.JWTMiddleware())

}
