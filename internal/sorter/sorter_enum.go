package sorter

type SorterType string

const (
	PerformanceSorterType SorterType = "performance"
	PriceSorterType       SorterType = "price"
	CreatedSorterType     SorterType = "created"
)
