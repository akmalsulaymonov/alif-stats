package stats

import (
	"fmt"

	"github.com/akmalsulaymonov/alif-bank/v2/pkg/types"
)

func Example_Stats() {
	payments := []types.Payment{
		{ID: 1111, Amount: 10_000_00, Category: "food", Status: "FAIL"},
		{ID: 2222, Amount: 3_000_00, Category: "car", Status: "OK"},
		{ID: 3333, Amount: 10_000_00, Category: "chemist", Status: "OK"},
		{ID: 4444, Amount: 7_000_00, Category: "food", Status: "INPROGRESS"},
	}

	avg := Avg(payments)
	total := TotalInCategory(payments, "food")
	fmt.Println(avg)
	fmt.Println(total)

	// output:
	// 500000
	// 700000
}
