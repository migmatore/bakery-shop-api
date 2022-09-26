package core

type OrderStatus struct {
	ID   int
	Name int
}

type GetOrderStatus struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type CreateOrderStatus struct {
}
