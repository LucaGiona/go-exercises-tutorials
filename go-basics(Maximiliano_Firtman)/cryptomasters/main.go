package main

import "frontendmasters.com/go/crypto/api"

func main(){

	rate, err := api.GetRate("BTC")

	print(rate,err)
}