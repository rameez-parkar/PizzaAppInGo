package Adapters

import (
	"Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMenu() Models.Menu {
	res, err := http.Get("http://localhost:9191/menu")
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
