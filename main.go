package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrKongKC/health-food/architecture"
	"github.com/mrKongKC/health-food/controller"
	"log"
)

func getAllHealthFoods(c echo.Context) error {
	y := c.QueryParam("year")
	status, value := controller.GetHealthFoodsHandler(y)
	return c.JSON(status, value)
}

func createHealthFood(c echo.Context) error {
	health := &architecture.HealthFood{}
	if err := c.Bind(health); err != nil {
		return c.JSON(400, err.Error())
	}
	status, value := controller.CreateHealthFoodHandler(*health)
	return c.JSON(status, value)
}

func main() {
	// initial database
	controller.InitializeDB()

	//create routing
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/health-foods", getAllHealthFoods)
	e.POST("/create-health-food", createHealthFood)

	// start server
	port := "5500"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}
