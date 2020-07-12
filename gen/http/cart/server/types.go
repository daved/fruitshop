// Code generated by goa v3.1.3, DO NOT EDIT.
//
// cart HTTP server types
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	cart "fruitshop/gen/cart"
	cartviews "fruitshop/gen/cart/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "cart" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// Name of the fruit
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of fruits
	Count *int `form:"Count,omitempty" json:"Count,omitempty" xml:"Count,omitempty"`
	// Cost of fruits
	CostPerItem *float64 `form:"CostPerItem,omitempty" json:"CostPerItem,omitempty" xml:"CostPerItem,omitempty"`
	// Total cost for the item
	TotalCost *float64 `form:"TotalCost,omitempty" json:"TotalCost,omitempty" xml:"TotalCost,omitempty"`
}

// CartManagementResponseCollection is the type of the "cart" service "get"
// endpoint HTTP response body.
type CartManagementResponseCollection []*CartManagementResponse

// CartManagementResponse is used to define fields on response body types.
type CartManagementResponse struct {
	// cartId is the unique id of the User.
	CartID string `form:"cartId" json:"cartId" xml:"cartId"`
	// Name of the fruit
	Name string `form:"Name" json:"Name" xml:"Name"`
	// Number of fruits
	Count int `form:"Count" json:"Count" xml:"Count"`
	// Cost of Each fruit
	CostPerItem int `form:"CostPerItem" json:"CostPerItem" xml:"CostPerItem"`
	// Total cost of fruits
	TotalCost int `form:"TotalCost" json:"TotalCost" xml:"TotalCost"`
}

// NewCartManagementResponseCollection builds the HTTP response body from the
// result of the "get" endpoint of the "cart" service.
func NewCartManagementResponseCollection(res cartviews.CartManagementCollectionView) CartManagementResponseCollection {
	body := make([]*CartManagementResponse, len(res))
	for i, val := range res {
		body[i] = marshalCartviewsCartManagementViewToCartManagementResponse(val)
	}
	return body
}

// NewAddPayload builds a cart service add endpoint payload.
func NewAddPayload(body *AddRequestBody, cartID string) *cart.AddPayload {
	v := &cart.AddPayload{
		Name:        *body.Name,
		Count:       *body.Count,
		CostPerItem: body.CostPerItem,
		TotalCost:   body.TotalCost,
	}
	v.CartID = cartID

	return v
}

// NewGetPayload builds a cart service get endpoint payload.
func NewGetPayload(cartID string) *cart.GetPayload {
	v := &cart.GetPayload{}
	v.CartID = cartID

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Name", "body"))
	}
	if body.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Count", "body"))
	}
	return
}
