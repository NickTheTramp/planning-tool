package handler

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", EmployeeHandler)
	e.POST("/employee", EmployeeHandler)
	e.POST("/employee/new", EmployeeFormHandler)
	e.GET("/employee/:id", EmployeeFormHandler)
	e.POST("/employee/:id", EmployeeFormHandler)
}
