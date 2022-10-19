package core

type Category struct {
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

//type CreateCategory struct {
//	Name        string `json:"name"`
//	Description string `json:"description,omitempty"`
//}
//
//type GetCategory struct {
//	ID          int    `json:"id"`
//	Name        string `json:"name"`
//	Description string `json:"description,omitempty"`
//}
