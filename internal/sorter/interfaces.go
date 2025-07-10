package sorter

import (
	"productSorter/internal/models"
)

type Sorter interface {
	Sort(products []*models.Product) []*models.Product
}
