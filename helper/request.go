package helper

type PaginationRequest struct {
	Page   int
	Size   int
	Filter Filter
	Order  Order
}

type Filter struct {
	Field string
	Value string
}

type Order struct {
	Direction string
	Field     string
}
