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

	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBites, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var cryptoRate datatypes.Rate
		err = json.Unmarshal(bodyBites, &cryptoRate)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: 20}
	return &rate, nil
}
