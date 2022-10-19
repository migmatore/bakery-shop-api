package core

type Product struct {
	ProductId         int     `json:"product_id"`
	Name              string  `json:"name"`
	Description       *string `json:"description,omitempty"`
	Price             float64 `json:"price"`
	ManufacturingDate string  `json:"manufacturing_date"`
	ExpirationDate    string  `json:"expiration_date"`
	CategoryId        *int    `json:"category_id,omitempty"`
	RecipeId          *int    `json:"recipe_id,omitempty"`
	ManufacturerId    int     `json:"manufacturer_id"`
}
