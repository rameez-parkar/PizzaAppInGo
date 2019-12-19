package main

import (
	"Services"
	"fmt"
)

var PizzaCount int

func main() {
	fmt.Println("####### PIZZA APP #######")
	fmt.Println("How many pizzas would you like to order?")
	fmt.Scanln(&PizzaCount)
	Services.GetMenu()
	for count := 0; count < PizzaCount; count++ {
		Services.ShowSizeOptions()
		Services.SelectSize()
		Services.ShowBaseOptions()
		Services.SelectBase()
		Services.ShowToppingOptions()
		Services.SelectToppings()
		Services.CalculatePizzaPrice()
		Services.AddPizza()
	}
	Services.DisplayCartPizzas()
	fmt.Println("\nProceed to Order? (y/n)")
	var orderConfirmation string
	fmt.Scanln(&orderConfirmation)
	if orderConfirmation == "y" {
		Services.ConfirmOrder()
	}
}
