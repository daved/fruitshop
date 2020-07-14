// Code generated by goa v3.1.3, DO NOT EDIT.
//
// discount HTTP client CLI support package
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	discount "fruitshop/gen/discount"
)

// BuildShowPayload builds the payload for the discount show endpoint from CLI
// flags.
func BuildShowPayload(discountShowUserID string) (*discount.ShowPayload, error) {
	var userID string
	{
		userID = discountShowUserID
	}
	v := &discount.ShowPayload{}
	v.UserID = userID

	return v, nil
}
