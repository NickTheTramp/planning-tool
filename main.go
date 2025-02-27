package main

import (
	"fmt"
	"github.com/NickTheTramp/planning-tool/handler"
	h "github.com/NickTheTramp/planning-tool/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	if err := h.InitializeDatabase(); err != nil {
		fmt.Println("Failed to initialize database")
		return
	}

	e := echo.New()

	handler.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
