// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payment HTTP client CLI support package
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	"encoding/json"
	"fmt"
	payment "fruitshop/gen/payment"
)

// BuildAddPayload builds the payload for the payment add endpoint from CLI
// flags.
func BuildAddPayload(paymentAddBody string, paymentAddUserID string) (*payment.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(paymentAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"ID\": \"Id earum explicabo tenetur sit.\",\n      \"amount\": 0.6465372050495056\n   }'")
		}
	}
	var userID string
	{
		userID = paymentAddUserID
	}
	v := &payment.AddPayload{
		ID:     body.ID,
		Amount: body.Amount,
	}
	v.UserID = userID

	return v, nil
}

// BuildGetPayload builds the payload for the payment get endpoint from CLI
// flags.
func BuildGetPayload(paymentGetBody string, paymentGetUserID string) (*payment.GetPayload, error) {
	var err error
	var body GetRequestBody
	{
		err = json.Unmarshal([]byte(paymentGetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"ID\": \"Quaerat atque.\"\n   }'")
		}
	}
	var userID string
	{
		userID = paymentGetUserID
	}
	v := &payment.GetPayload{
		ID: body.ID,
	}
	v.UserID = userID

	return v, nil
}
