package core

import (
	"github.com/jackc/pgtype"
	"time"
)

type Product struct {
	ProductId         int         `json:"product_id"`
	Name              string      `json:"name"`
	ImagePath         string      `json:"image_path"`
	Description       *string     `json:"description,omitempty"`
	Price             float32     `json:"price"`
	ManufacturingDate pgtype.Date `json:"manufacturing_date"`
	ExpirationDate    pgtype.Date `json:"expiration_date"`
	CategoryId        *int        `json:"category_id,omitempty"`
	RecipeId          *int        `json:"recipe_id,omitempty"`
	StoreId           int         `json:"store_id"`
	UnitStock         int         `json:"unit_stock"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
}

type PatchProductDTO struct {
	Name              *string  `json:"name,omitempty"`
	ImagePath         *string  `json:"image_path,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Price             *float32 `json:"price,omitempty"`
	ManufacturingDate *string  `json:"manufacturing_date,omitempty"`
	ExpirationDate    *string  `json:"expiration_date,omitempty"`
	CategoryId        *int     `json:"category_id,omitempty"`
	RecipeId          *int     `json:"recipe_id,omitempty"`
	UnitStock         *int     `json:"unit_stock,omitempty"`
}

type PatchProduct struct {
	Name              *string   `json:"name,omitempty"`
	ImagePath         *string   `json:"image_path,omitempty"`
	Description       *string   `json:"description,omitempty"`
	Price             *float32  `json:"price,omitempty"`
	ManufacturingDate *string   `json:"manufacturing_date,omitempty"`
	ExpirationDate    *string   `json:"expiration_date,omitempty"`
	CategoryId        *int      `json:"category_id,omitempty"`
	RecipeId          *int      `json:"recipe_id,omitempty"`
	UnitStock         *int      `json:"unit_stock,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type CreateProductDTO struct {
	Name              string  `json:"name"`
	ImagePath         string  `json:"image_path"`
	Description       *string `json:"description,omitempty"`
	Price             float32 `json:"price"`
	ManufacturingDate string  `json:"manufacturing_date"`
	ExpirationDate    string  `json:"expiration_date"`
	CategoryId        *int    `json:"category_id,omitempty"`
	RecipeId          *int    `json:"recipe_id,omitempty"`
	UnitStock         int     `json:"unit_stock"`
}

type CreateProduct struct {
	Name              string    `json:"name"`
	ImagePath         string    `json:"image_path"`
	Description       *string   `json:"description,omitempty"`
	Price             float32   `json:"price"`
	ManufacturingDate string    `json:"manufacturing_date"`
	ExpirationDate    string    `json:"expiration_date"`
	CategoryId        *int      `json:"category_id,omitempty"`
	RecipeId          *int      `json:"recipe_id,omitempty"`
	StoreId           int       `json:"store_id"`
	UnitStock         int       `json:"unit_stock"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

func NewCreateProductFromDTO(dto *CreateProductDTO, storeId int) *CreateProduct {
	return &CreateProduct{
		Name:              dto.Name,
		ImagePath:         dto.ImagePath,
		Description:       dto.Description,
		Price:             dto.Price,
		ManufacturingDate: dto.ManufacturingDate,
		ExpirationDate:    dto.ExpirationDate,
		CategoryId:        dto.CategoryId,
		RecipeId:          dto.RecipeId,
		StoreId:           storeId,
		UnitStock:         dto.UnitStock,
		CreatedAt:         time.Now(),
	}
}

func NewPatchProductFromDTO(dto *PatchProductDTO) *PatchProduct {
	return &PatchProduct{
		Name:              dto.Name,
		ImagePath:         dto.ImagePath,
		Description:       dto.Description,
		Price:             dto.Price,
		ManufacturingDate: dto.ManufacturingDate,
		ExpirationDate:    dto.ExpirationDate,
		CategoryId:        dto.CategoryId,
		RecipeId:          dto.RecipeId,
		UnitStock:         dto.UnitStock,
		UpdatedAt:         time.Now(),
	}
}
