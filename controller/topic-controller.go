package controller

import (
	"github.com/mrKongKC/health-food/services"
)

func InitializeDB() {
	services.ConnectDB()
}

func GetAllHeathFoods(y string) (string, any) {
	if y == "" {
		status, value := services.GetAllHeathFoods()
		return status, value
	}
	return "not", 0
}
