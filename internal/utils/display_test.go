package utils_test

import (
	"bytes"
	"os"
	"productSorter/internal/models"
	"productSorter/internal/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayProductTable(t *testing.T) {
	products := []*models.Product{
		{ID: 1, Name: "Table", Price: 100.0, Created: "2022-01-01", SalesCount: 10, ViewsCount: 100},
	}

	// Capture stdout
	var buf bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	utils.DisplayProductTable(products)
	_ = w.Close()
	_, _ = buf.ReadFrom(r)
	os.Stdout = originalStdout

	output := strings.ReplaceAll(buf.String(), "\t", " ")

	assert.Contains(t, output, "S/N")
	assert.Contains(t, output, "Table")
	assert.Contains(t, output, "0.1000")

	assert.Contains(t, output, "1")
}
