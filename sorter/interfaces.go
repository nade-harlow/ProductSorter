package sorter

import "productSorter/models"

type Sorter interface {
	Sort(products []*models.Product) []*models.Product
}
