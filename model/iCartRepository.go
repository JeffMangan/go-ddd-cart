package model

type ICartRepository interface {
	AddItemToCart() error
	RemoveItemFromCart() error
	SaveCartForLater() error
	DeleteCart() error
}