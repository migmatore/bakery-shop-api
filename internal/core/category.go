package core

type Category struct {
	ID   int
	Name string
}

type CreateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type GetCategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
