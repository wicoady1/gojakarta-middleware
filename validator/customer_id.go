package validator

import "net/http"

func ValidCustomerID(r *http.Request) bool {
	return true
}
