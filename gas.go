package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	GasURL = "https://www.ethgasstationapi.com/api/%s"
)

type GasPrices struct {
	Low      string `json:"low"`
	Standard string `json:"standard"`
	Fast     string `json:"fast"`
}

// GetGasPrice retrieves the current gas rates
func GetGasPrice(rate string) (string, error) {

	var price string

	reqUrl := fmt.Sprintf(GasURL, rate)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return price, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return price, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return price, err
	}

	resultsString := string(results)

	price = strings.Split(resultsString, ".")[0]

	return price, nil
}

func GetGasPrices() (GasPrices, error) {
	var gasPrices GasPrices

	low, err := GetGasPrice("low")
	standard, err := GetGasPrice("standard")
	fast, err := GetGasPrice("fast")

	if err != nil {
		return gasPrices, err
	}

	gasPrices = GasPrices{
		Low:      low,
		Standard: standard,
		Fast:     fast,
	}

	return gasPrices, nil
}
