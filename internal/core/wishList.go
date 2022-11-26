package core

type WishList struct {
	WishListId int `json:"wish_list_id"`
}

type WishListItem struct {
	WishListItemId int `json:"wish_list_item_id"`
	WishListId     int `json:"wish_list_id"`
}
