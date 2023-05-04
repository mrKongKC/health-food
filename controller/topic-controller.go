package controller

import (
	"strconv"

	"github.com/mrKongKC/health-food/architecture"
	"github.com/mrKongKC/health-food/services"
)

func InitializeDB() {
	services.ConnectDB()
}

func GetHealthFoodsHandler(y string) (int, any) {
	var status int
	var value any
	if y == "" {
		status, value = services.GetAllHealthFoods()
		return status, value
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return 400, err.Error()
	} else {
		status, value = services.QueryHealthFoodsByYear(year)
		return status, value
	}

}

func CreateHealthFoodHandler(health architecture.HealthFood) (int, any) {
	status, value := services.CreateHealthFood(health)
	return status, value
}
