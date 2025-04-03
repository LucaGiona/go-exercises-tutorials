package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3characters minimum required: %d received", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBites, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBites, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: float64(response.Bid)}
	return &rate, nil
}
