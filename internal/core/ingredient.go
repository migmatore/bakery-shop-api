package core

type Ingredient struct {
	IngredientId      int     `json:"ingredient_id"`
	Name              string  `json:"name"`
	Description       *string `json:"description,omitempty"`
	RemainingQuantity int     `json:"remaining_quantity"`
	WeightUnitId      int     `json:"weight_unit"`
	SupplierId        int     `json:"supplier_id"`
}
