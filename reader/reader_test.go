package reader_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"productSorter/models"
	"productSorter/reader"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProducts(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(t *testing.T, dir string)
		expectError bool
		expectedLen int
		errorSubstr string
	}{

		{
			name: "valid file",
			setup: func(t *testing.T, dir string) {
				data := []*models.Product{
					{ID: 1, Name: "Table", Price: 10.0, Created: "2022-01-01", SalesCount: 10, ViewsCount: 100},
				}
				_ = os.MkdirAll(filepath.Join(dir, "data"), 0755)
				filePath := filepath.Join(dir, "data", "products.json")
				bytes, _ := json.Marshal(data)
				_ = os.WriteFile(filePath, bytes, 0644)
			},
			expectError: false,
			expectedLen: 1,
		},
		{
			name: "missing file",
			setup: func(t *testing.T, dir string) {
				_ = os.MkdirAll(filepath.Join(dir, "data"), 0755)
				// intentionally do not create products.json
			},
			expectError: true,
			errorSubstr: "failed to read",
		},
		{
			name: "invalid JSON",
			setup: func(t *testing.T, dir string) {
				_ = os.MkdirAll(filepath.Join(dir, "data"), 0755)
				filePath := filepath.Join(dir, "data", "products.json")
				_ = os.WriteFile(filePath, []byte("not-json"), 0644)
			},
			expectError: true,
			errorSubstr: "failed to parse",
		},
		{
			name: "empty list",
			setup: func(t *testing.T, dir string) {
				_ = os.MkdirAll(filepath.Join(dir, "data"), 0755)
				filePath := filepath.Join(dir, "data", "products.json")
				_ = os.WriteFile(filePath, []byte("[]"), 0644)
			},
			expectError: true,
			errorSubstr: "no products found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			_ = os.Chdir(tmpDir)

			tc.setup(t, tmpDir)

			products, err := reader.ReadProducts()

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorSubstr != "" {
					assert.Contains(t, err.Error(), tc.errorSubstr)
				}
				assert.Nil(t, products)
			} else {
				assert.NoError(t, err)
				assert.Len(t, products, tc.expectedLen)
			}
		})
	}
}
