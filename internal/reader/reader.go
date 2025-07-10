package reader

import (
	"encoding/json"
	"fmt"
	"os"
	"productSorter/internal/models"
)

const productFilePath = "data/products.json"

func ReadProducts() ([]*models.Product, error) {
	var products []*models.Product

	file, err := os.ReadFile(productFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read .products.json: %w", err)
	}

	if err = json.Unmarshal(file, &products); err != nil {
		return nil, fmt.Errorf("failed to parse .products.json: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found in .products.json")
	}

	return products, nil
}
