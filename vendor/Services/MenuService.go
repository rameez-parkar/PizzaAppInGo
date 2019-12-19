package Services

import (
	"Adapters"
	"Models"
	"fmt"
	"strconv"
)

var menu Models.Menu

func GetMenu() {
	menu = Adapters.GetMenu()
}

func ShowBaseOptions() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	fmt.Println("Choose Pizza Base (Enter Base Number) : ")
	for base := 0; base < len(menu.Bases); base++ {
		fmt.Println("Base Id : " + strconv.Itoa(menu.Bases[base].Id))
		fmt.Println("Name : " + menu.Bases[base].Name)
		fmt.Println("Price : " + strconv.FormatFloat(float64(menu.Bases[base].Price), 'f', 2, 64) + "\n")
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
}

func ShowSizeOptions() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	fmt.Println("Choose Pizza Size (Enter Size as S, M or L) : ")
	for size := 0; size < len(menu.Sizes); size++ {
		fmt.Println("Size Id : " + menu.Sizes[size].Id)
		fmt.Println("Name : " + menu.Sizes[size].Size)
		fmt.Println("Price : " + strconv.FormatFloat(float64(menu.Sizes[size].Price), 'f', 2, 64) + "\n")
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
}

func ShowToppingOptions() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	fmt.Println("Choose Toppings (Enter Toppings as 0,5,6 ) : ")
	for topping := 0; topping < len(menu.Toppings); topping++ {
		fmt.Println("Topping Id : " + strconv.Itoa(menu.Toppings[topping].Id))
		fmt.Println("Name : " + menu.Toppings[topping].Name)
		fmt.Println("Price : " + strconv.FormatFloat(float64(menu.Toppings[topping].Price), 'f', 2, 64) + "\n")
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
}
