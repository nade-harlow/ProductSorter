package sorter

import (
	"productSorter/internal/models"
	"sort"
)

type PriceSorter struct{}

func (ps PriceSorter) Sort(products []*models.Product) []*models.Product {
	sorted := make([]*models.Product, len(products))
	copy(sorted, products)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Price < sorted[j].Price
	})
	return sorted
}
