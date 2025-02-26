package handler

import "net/http"

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", EmployeeHandler)
	mux.HandleFunc("/employee/{id}", EmployeeFormHandler)

	return mux
}
