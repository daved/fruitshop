// Code generated by goa v3.2.0, DO NOT EDIT.
//
// HTTP request path constructors for the coupon service.
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	"fmt"
)

// AddCouponPath returns the URL path to the coupon service add HTTP endpoint.
func AddCouponPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/coup/%v", userID)
}
