package core

import "github.com/jackc/pgtype"

type Product struct {
	ProductId         int         `json:"product_id"`
	Name              string      `json:"name"`
	ImagePath         string      `json:"image_path"`
	Description       *string     `json:"description,omitempty"`
	Price             float64     `json:"price"`
	ManufacturingDate pgtype.Date `json:"manufacturing_date"`
	ExpirationDate    pgtype.Date `json:"expiration_date"`
	CategoryId        *int        `json:"category_id,omitempty"`
	RecipeId          *int        `json:"recipe_id,omitempty"`
	ManufacturerId    int         `json:"manufacturer_id"`
}

type PatchProduct struct {
	Name              *string      `json:"name,omitempty"`
	ImagePath         *string      `json:"image_path,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Price             *float64     `json:"price,omitempty"`
	ManufacturingDate *pgtype.Date `json:"manufacturing_date,omitempty"`
	ExpirationDate    *pgtype.Date `json:"expiration_date,omitempty"`
	CategoryId        *int         `json:"category_id,omitempty"`
	RecipeId          *int         `json:"recipe_id,omitempty"`
	ManufacturerId    *int         `json:"manufacturer_id,omitempty"`
}
