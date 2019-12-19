package Services

import (
	"Models"
	"fmt"
	"strconv"
)

var Cart Models.Cart

func AddPizza() {
	Cart.Pizzas = append(Cart.Pizzas, Pizza)
}

func DisplayCartPizzas() {
	fmt.Println("\n-----------------CART ITEMS-----------------")
	for pizza := 0; pizza < len(Cart.Pizzas); pizza++ {
		fmt.Println("\nPizza Size : " + Cart.Pizzas[pizza].Size.Size)
		fmt.Println("Pizza Base : " + Cart.Pizzas[pizza].Base.Name)
		fmt.Println("Pizza Toppings : ")
		for topping := 0; topping < len(Cart.Pizzas[pizza].Toppings); topping++ {
			fmt.Print(Cart.Pizzas[pizza].Toppings[topping].Name + ", ")
		}
		fmt.Println("\nPizza Price : " + strconv.FormatFloat(Cart.Pizzas[pizza].Price, 'f', 6, 64))
	}
	fmt.Println("\nFINAL AMOUNT : " + strconv.FormatFloat(CalculateTotalPrice(), 'f', 6, 64))
	fmt.Println("--------------------------------------------")
}

func CalculateTotalPrice() float64 {
	var totalPrice float64
	for pizza := 0; pizza < len(Cart.Pizzas); pizza++ {
		totalPrice = totalPrice + Cart.Pizzas[pizza].Price
	}
	return totalPrice
}
