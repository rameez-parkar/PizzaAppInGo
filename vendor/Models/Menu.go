package Models

type Menu struct {
	Bases    []Base    `json:"base"`
	Sizes    []Size    `json:"size"`
	Toppings []Topping `json:"topping"`
}
