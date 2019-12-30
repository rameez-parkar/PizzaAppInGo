package Services

import (
	"Models"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Order Models.Order

func ConfirmOrder() {
	Order.OrderId = rand.Intn(500)
	Order.Cart = Cart
	fmt.Println("\n------------------ORDER DETAILS--------------------------")
	fmt.Println("Order Successful!")
	fmt.Println("ORDER ID : " + strconv.Itoa(Order.OrderId))
	dateTime := time.Now()
	Order.Date = dateTime.Format("2006-01-02")
	Order.Time = dateTime.Format("15:04:05")
	fmt.Print("Date : ")
	fmt.Println(Order.Date)
	fmt.Print("Time : ")
	fmt.Println(Order.Time)
	DisplayCartPizzas()
	fmt.Println("--------------------------------------------")
	http.HandleFunc("/", HostOrder)
	http.ListenAndServe(":9080", nil)
}

func HostOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
