package Models

type Pizza struct {
	Base     Base
	Size     Size
	Toppings []Topping
	Price float64
}
