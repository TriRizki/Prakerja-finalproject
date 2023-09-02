package routes

import (
	"Prakerja-finalproject/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	
    e.GET("/pegawai", controllers.FetchAllPegawai)
	e.GET("/pegawai/:id", controllers.FetchById)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai/:id", controllers.DeletePegawai)

	return e
}