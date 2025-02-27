package handler

import (
	h "github.com/NickTheTramp/planning-tool/helper"
	"github.com/NickTheTramp/planning-tool/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func EmployeeFormHandler(c echo.Context) error {
	route := h.Route{
		Title:    "Employee Form",
		FileName: "employeeForm.html",
	}

	id := r.PathValue("id")
	log.Println(id)

	employee := model.Employee{}

	if r.Method == http.MethodPost {
		employee.FirstName = r.FormValue("first_name")
		employee.LastName = r.FormValue("last_name")

		if result := h.DB.Create(&employee); result.Error != nil {
			log.Println("Failed to insert data:", result.Error)

			return
		}
	}

	h.RenderTemplate(c.Response(), route, employee)
}

func EmployeeHandler(c echo.Context) error {
	route := h.Route{
		Title:    "Employee Form",
		FileName: "employeeForm.html",
	}

	employee := model.Employee{}

	if r.Method == http.MethodPost {
		employee.FirstName = r.FormValue("first_name")
		employee.LastName = r.FormValue("last_name")

		if result := h.DB.Create(&employee); result.Error != nil {
			log.Println("Failed to insert data:", result.Error)

			return
		}
	}

	h.RenderTemplate(c.Response(), route, employee)
}
