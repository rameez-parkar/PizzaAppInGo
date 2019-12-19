package Services

import (
	"Models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Pizza Models.Pizza

func SelectBase() {
	fmt.Print("Your Base choice : ")
	var baseChoice int
	fmt.Scanf("%d\n", &baseChoice)
	base := ValidateBaseInput(baseChoice)
	if (base != Models.Base{}) {
		Pizza.Base = base
	}
}

func SelectSize() {
	fmt.Print("Your Size choice : ")
	var sizeChoice string
	fmt.Scanln(&sizeChoice)
	sizeChoice = strings.ToUpper(sizeChoice)
	size := ValidateSizeInput(sizeChoice)
	if (size != Models.Size{}) {
		Pizza.Size = size
	}
}

func SelectToppings() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Toppings choice : ")
	var toppingsChoiceInput string
	toppingsChoiceInput, _ = reader.ReadString('\n')
	toppingsChoiceInput = strings.ReplaceAll(toppingsChoiceInput, " ", "")
	toppingsChoiceInput = strings.TrimSuffix(toppingsChoiceInput, "\n")
	fmt.Println(toppingsChoiceInput)
	toppingsChoiceString := strings.Split(toppingsChoiceInput, ",")
	fmt.Println(toppingsChoiceString)
	var toppingsChoice = make([]int, len(toppingsChoiceString))
	for i := 0; i < len(toppingsChoiceString); i++ {
		value, _ := strconv.Atoi(toppingsChoiceString[i])
		toppingsChoice[i] = value
		fmt.Println(toppingsChoice)
	}
	toppings := ValidateToppingsInput(toppingsChoice)
	fmt.Println(toppings)
	Pizza.Toppings = toppings
}

func CalculatePizzaPrice() {
	var pizzaPrice float64
	pizzaPrice = pizzaPrice + Pizza.Size.Price
	pizzaPrice = pizzaPrice + Pizza.Base.Price
	for topping := 0; topping < len(Pizza.Toppings); topping++ {
		pizzaPrice = pizzaPrice + Pizza.Toppings[topping].Price
	}
	Pizza.Price = pizzaPrice
}
