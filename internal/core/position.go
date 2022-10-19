package core

type Position struct {
	PositionId  int     `json:"position_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
