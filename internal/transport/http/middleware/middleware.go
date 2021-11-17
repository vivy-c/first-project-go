package middleware

import (
	"github.com/labstack/echo"
)

type Middleware struct {
}

//instance (objek) buat dipanggil di kelas lain
func NewMidleware() *Middleware {
	return &Middleware{}
}

//disini ngeluarin http response writer sama request

func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*") //* bisa buat semua ip, kalo ada yg dibok bisa ditulis dsini
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization,Origin,Accept,datetime,signature,Content-Type")
		c.Response().Header().Set("Content-Type", "application/json") //cma bisa nrima json
		return next(c)
	}
}