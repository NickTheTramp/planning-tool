package main

import (
	"fmt"
	h "github.com/NickTheTramp/planning-tool/helper"
	"github.com/NickTheTramp/planning-tool/model"
	"log"
	"net/http"
)

func main() {
	if err := h.InitializeDatabase(); err != nil {
		fmt.Println("Failed to initialize database")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
				http.Error(w, "Failed to save employee", http.StatusInternalServerError)
				return
			}
		}

		h.ViewTemplate(w, route, employee)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to listen and serve")
		return
	}
}
