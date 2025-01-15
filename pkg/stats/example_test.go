package stats

import (
	"fmt"

	"github.com/akmalsulaymonov/alif-bank/v2/pkg/types"
)

func Example_Stats() {
	payments := []types.Payment{
		{ID: 1111, Amount: 10_000_00, Category: "food"},
		{ID: 2222, Amount: 3_000_00, Category: "car"},
		{ID: 3333, Amount: 10_000_00, Category: "chemist"},
		{ID: 4444, Amount: 7_000_00, Category: "food"},
	}

	avg := Avg(payments)
	total := TotalInCategory(payments, "food")
	fmt.Println(avg)
	fmt.Println(total)

	// output:
	// 750000
	// 1700000
}
