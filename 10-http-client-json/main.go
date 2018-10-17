package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IfConfigResponse struct {
	IP         string  `json:"ip"`
	IPDecimal  int     `json:"ip_decimal"`
	Country    string  `json:"country"`
	CountryEu  bool    `json:"country_eu"`
	CountryIso string  `json:"country_iso"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func main() {
	response, err := http.Get("https://ifconfig.co/json")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	ifConfigResponse := &IfConfigResponse{}
	json.Unmarshal(body, ifConfigResponse)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("City: %s\n", ifConfigResponse.City)
	fmt.Printf("Country: %s\n", ifConfigResponse.Country)
	fmt.Printf("IP: %s\n", ifConfigResponse.IP)
}
