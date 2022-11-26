package core

type Cart struct {
	CartId       int     `json:"cart_id"`
	ProductCount int     `json:"product_count"`
	TotalPrice   float32 `json:"total_price"`
}

type CartItem struct {
	CartItemId int     `json:"cart_item_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float32 `json:"price"`
	CartId     int     `json:"cart_id"`
}
