package main

import (
	"fmt"
	"github.com/NickTheTramp/planning-tool/handler"
	h "github.com/NickTheTramp/planning-tool/helper"
	"net/http"
)

func main() {
	if err := h.InitializeDatabase(); err != nil {
		fmt.Println("Failed to initialize database")
		return
	}

	mux := handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to listen and serve")
		return
	}
}
