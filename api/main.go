package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Coffee struct {
	Country  string `json:"country"`
	Region   string `json:"region"`
	Producer string `json:"producer"`
	Name     string `json:"name"`
	Process  string `json:"process"`
	Flavours string `json:"flavours"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData())
}

func dataReader() []byte {
	data, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func jsonData() []byte {
	data := dataReader()
	var coffees []Coffee
	err := json.Unmarshal(data, &coffees)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(coffees)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}
