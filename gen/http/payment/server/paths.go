// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the payment service.
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	"fmt"
)

// AddPaymentPath returns the URL path to the payment service add HTTP endpoint.
func AddPaymentPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/payment/pay/%v", userID)
}

// GetPaymentPath returns the URL path to the payment service get HTTP endpoint.
func GetPaymentPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/payment/%v", userID)
}
