package stats

import "github.com/akmalsulaymonov/alif-bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, item := range payments {
		sum += item.Amount
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, item := range payments {
		if item.Category == category {
			sum += item.Amount
		}
	}
	return sum
}
