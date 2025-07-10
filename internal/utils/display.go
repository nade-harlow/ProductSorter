package utils

import (
	"fmt"
	"os"
	"productSorter/internal/models"
	"text/tabwriter"
)

func DisplayProductTable(products []*models.Product) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "S/N\tID\tName\tPrice\tCreated\tSales\tViews\tSales/View")

	for sn, p := range products {
		ratio := float64(p.SalesCount) / float64(p.ViewsCount)
		fmt.Fprintf(w, "%d\t%d\t%s\t%.2f\t%s\t%d\t%d\t%.4f\n",
			sn+1, p.ID, p.Name, p.Price, p.Created, p.SalesCount, p.ViewsCount, ratio)
	}

	w.Flush()
}
