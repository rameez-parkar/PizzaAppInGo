package Services

import (
	"Models"
)

func ValidateBaseInput(baseChoice int) Models.Base {
	for base := 0; base < len(menu.Bases); base++ {
		if menu.Bases[base].Id == baseChoice {
			return menu.Bases[base]
		}
	}
	return Models.Base{}
}

func ValidateSizeInput(sizeChoice string) Models.Size {
	for size := 0; size < len(menu.Sizes); size++ {
		if menu.Sizes[size].Id == sizeChoice {
			return menu.Sizes[size]
		}
	}
	return Models.Size{}
}

func ValidateToppingsInput(toppingsChoices []int) []Models.Topping {
	var toppings []Models.Topping
	for toppingsChoice := 0; toppingsChoice < len(toppingsChoices); toppingsChoice++ {
		for topping := 0; topping < len(menu.Toppings); topping++ {
			if menu.Toppings[topping].Id == toppingsChoices[toppingsChoice] {
				toppings = append(toppings, menu.Toppings[topping])
				break
			}
		}
		// if toppingsChoice != len(toppings) {
		// 	return nil
		// }
	}
	return toppings
}
