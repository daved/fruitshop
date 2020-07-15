// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the discount service.
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	"fmt"
)

// GetDiscountPath returns the URL path to the discount service get HTTP endpoint.
func GetDiscountPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/discount/%v", userID)
}
