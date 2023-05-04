package services

import (
	"fmt"
	"github.com/mrKongKC/health-food/architecture"
	"github.com/mrKongKC/health-food/db"
	"log"
)

func ConnectDB() {
	db.Conn()
}

func CreateHealthFood(health architecture.HealthFood) (int, any) {
	stmt, err := db.Db.Prepare(`
	INSERT INTO HealthFood( title, ingredient, recipe, year, isApprove)
	VALUES (?,?,?,?,?);
	`)
	if err != nil {
		return 500, err.Error()
	}
	defer stmt.Close()

	b := fmt.Sprintf("%v", health.IsApprove)
	result, err := stmt.Exec(health.Title, health.Ingredient, health.Recipe, health.Year, b)
	
	switch {
	case err == nil:
		id, _ := result.LastInsertId()
		health.HealthFoodId = id
		return 201, health
	case err.Error() == "UNIQUE constraint violation":
		return 409, "movie already exists"
	default:
		return 500, err.Error()
	}
}

func GetAllHealthFoods() (int, any) {
	hf := []architecture.HealthFood{}
	rows, err := db.Db.Query(`SELECT healthFoodId, title, ingredient, recipe, year, isApprove
		FROM HealthFood`)
	if err != nil {
		log.Fatal("query error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var health architecture.HealthFood
		if err := rows.Scan(&health.HealthFoodId, &health.Title, &health.Ingredient, &health.Recipe, &health.Year, &health.IsApprove); err != nil {
			return 500, err.Error()
		}
		hf = append(hf, health)
	}

	if err := rows.Err(); err != nil {
		return 500, err.Error()
	}

	return 200, hf
}

func QueryHealthFoodsByYear(year int) (int, any) {
	hf := []architecture.HealthFood{}
	rows, err := db.Db.Query(`SELECT healthFoodId, title, ingredient, recipe, year, isApprove
	FROM HealthFood
	WHERE year = ?`, year)
	if err != nil {
		return 500, err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var health architecture.HealthFood
		if err := rows.Scan(&health.HealthFoodId, &health.Title, &health.Ingredient, &health.Recipe, &health.Year, &health.IsApprove); err != nil {
			return 500, err.Error()
		}
		hf = append(hf, health)
	}

	if err := rows.Err(); err != nil {
		return 500, err.Error()
	}

	return 200, hf
}
