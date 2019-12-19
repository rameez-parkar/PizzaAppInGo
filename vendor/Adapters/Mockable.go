package Adapters

import (
	"Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMenu() Models.Menu {
	res, err := http.Get("https://demo6059093.mockable.io/pizzaapp")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var menu Models.Menu
	json.Unmarshal(data, &menu)
	res.Body.Close()
	return menu
}
