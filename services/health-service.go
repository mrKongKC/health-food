package services

import (
	"github.com/mrKongKC/health-food/db"
	"log"
)

var database = db.Db

type HealthFood struct {
	HealthFoodId int64  `json:"healthFoodId"`
	Title        string `json:"title"`
	Ingredient   string `json:"ingredient"`
	Recipe       string `json:recipe`
	Year         int    `json:"year"`
	IsApprove    bool   `json:"isApprove"`
}

func ConnectDB() {
	db.Conn()
}

func GetAllHeathFoods() (string, any) {
	hf := []HealthFood{}

	rows, err := database.Query(`SELECT healthFoodId, title, ingredient, recipe, year, isApprove
		FROM HealthFood`)
	if err != nil {
		log.Fatal("query error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var health HealthFood
		if err := rows.Scan(&health.HealthFoodId, &health.Title, &health.Ingredient, &health.Recipe, &health.Year, &health.IsApprove); err != nil {
			return "InternalServerError", err.Error()
		}
		hf = append(hf, health)
	}

	if err := rows.Err(); err != nil {
		return "InternalServerError", err.Error()
	}

	return "OK", hf
}
