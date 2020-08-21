package controller

import (
	"bookstore/pkg/api"
	"github.com/labstack/echo/v4"
)

func BookController(g *echo.Group) {

	g.POST("/api/customer", api.Createcustomer)
	g.PUT("/api/customer", api.Updatecustomer)
}
