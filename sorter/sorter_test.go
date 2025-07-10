package sorter_test

import (
	"testing"

	"productSorter/models"
	"productSorter/sorter"

	"github.com/stretchr/testify/assert"
)

func sampleProducts() []*models.Product {
	return []*models.Product{
		{ID: 1, Name: "A", Price: 100.00, SalesCount: 10, ViewsCount: 100},
		{ID: 2, Name: "B", Price: 50.00, SalesCount: 20, ViewsCount: 100},
		{ID: 3, Name: "C", Price: 75.00, SalesCount: 5, ViewsCount: 100},
	}
}

func TestPriceSorter(t *testing.T) {
	ps := sorter.PriceSorter{}
	sorted := ps.Sort(sampleProducts())

	assert.Equal(t, 2, sorted[0].ID)
	assert.Equal(t, 3, sorted[1].ID)
	assert.Equal(t, 1, sorted[2].ID)
}

func TestPerformanceSorter(t *testing.T) {
	ps := sorter.PerformanceSorter{}
	sorted := ps.Sort(sampleProducts())

	assert.Equal(t, 2, sorted[0].ID)
	assert.Equal(t, 1, sorted[1].ID)
	assert.Equal(t, 3, sorted[2].ID)
}

func TestRegistry(t *testing.T) {
	reg := sorter.NewRegistry()

	ps := sorter.PriceSorter{}
	reg.Register(sorter.PriceSorterType, ps)

	got := reg.Get(sorter.PriceSorterType)
	assert.NotNil(t, got)

	sorted := got.Sort(sampleProducts())
	assert.Equal(t, 2, sorted[0].ID)
}
