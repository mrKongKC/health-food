package architecture

type HealthFood struct {
	HealthFoodId int64  `json:"healthFoodId"`
	Title        string `json:"title"`
	Ingredient   string `json:"ingredient"`
	Recipe       string `json:"recipe"`
	Year         int    `json:"year"`
	IsApprove    bool   `json:"isApprove"`
}



