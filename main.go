package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrKongKC/health-food/controller"
	"log"
	"net/http"
)

func getAllHeathFoods(c echo.Context) error {
	y := c.QueryParam("year")
	staus, value := controller.GetAllHeathFoods(y)
	if staus == "InternalServerError" {
		return c.JSON(http.StatusInternalServerError, value)
	} else {
		return c.JSON(http.StatusOK, value)
	}

}

func main() {
	controller.InitializeDB()
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/health-foods", getAllHeathFoods)
	port := "5500"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}
