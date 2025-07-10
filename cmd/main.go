package main

import (
	"flag"
	"fmt"
	"os"
	"productSorter/internal/reader"
	"productSorter/internal/sorter"
	"productSorter/internal/utils"
)

func main() {
	sortKey := flag.String("sort", string(sorter.PerformanceSorterType), "sorting strategy: performance | price")
	flag.Parse()

	products, err := reader.ReadProducts()
	if err != nil {
		fmt.Printf("Error reading products: %v\n", err)
		return
	}

	registry := sorter.NewRegistry()
	registry.Register(sorter.PriceSorterType, sorter.PriceSorter{})
	registry.Register(sorter.PerformanceSorterType, sorter.PerformanceSorter{})

	s := registry.Get(sorter.SorterType(*sortKey))
	if s == nil {
		fmt.Fprintf(os.Stderr, "Unknown sorter type: %s\n", *sortKey)
		fmt.Fprintln(os.Stderr, "Available: performance, price")
		os.Exit(1)
	}

	sorted := s.Sort(products)
	utils.DisplayProductTable(sorted)
}
