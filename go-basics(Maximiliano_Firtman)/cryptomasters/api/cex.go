package api

import (
	"fmt"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (datatypes.Rate, error){
	
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprint(apiUrl, upCurrency))

}