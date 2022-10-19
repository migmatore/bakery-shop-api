package core

type RecipeIngredient struct {
	RecipeIngredientId int `json:"recipe_ingredient_id"`
	RecipeId           int `json:"recipe_id"`
	IngredientId       int `json:"ingredient_id"`
	Quantity           int `json:"quantity"`
}
