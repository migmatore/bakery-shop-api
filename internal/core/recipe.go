package core

type Recipe struct {
	RecipeId    int    `json:"recipe_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}
