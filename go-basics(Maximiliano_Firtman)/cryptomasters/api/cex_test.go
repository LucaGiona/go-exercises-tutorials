package api_test

import "frontendmasters.com/go/crypto/api"
import "testing"

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}