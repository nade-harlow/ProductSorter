package models

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Created    string  `json:"created"`
	SalesCount int     `json:"sales_count"`
	ViewsCount int     `json:"views_count"`
}
