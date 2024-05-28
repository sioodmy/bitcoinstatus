package price

import (
	"encoding/json"
	"fmt"
	"math"

	"net/http"
	"strconv"
)

func GetPrice() (string, error) {

	url := "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	type BinanceApi struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	var priceResponse BinanceApi
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return "", err
	}

	price := RoundPrice(priceResponse.Price)
	formatted := fmt.Sprintf("%v$", price)

	return formatted, nil
}

func RoundPrice(price string) float64 {
	parsed, _ := strconv.ParseFloat(price, 32)

	return math.Floor(parsed)
}
