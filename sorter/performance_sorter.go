package sorter

import (
	"productSorter/models"
	"sort"
)

type PerformanceSorter struct{}

func (ps PerformanceSorter) Sort(products []*models.Product) []*models.Product {
	sorted := make([]*models.Product, len(products))
	copy(sorted, products)
	sort.Slice(sorted, func(i, j int) bool {
		ri := float64(sorted[i].SalesCount) / float64(sorted[i].ViewsCount)
		rj := float64(sorted[j].SalesCount) / float64(sorted[j].ViewsCount)
		return ri > rj
	})
	return sorted
}
