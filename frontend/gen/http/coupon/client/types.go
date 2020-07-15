// Code generated by goa v3.2.0, DO NOT EDIT.
//
// coupon HTTP client types
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	couponviews "fruitshop/frontend/gen/coupon/views"
)

// AddResponseBody is the type of the "coupon" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// ID is the unique id of the User.
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
	// ID is the unique id of the Users coupon.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Name of the coupon.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status of  Users coupon.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Users coupon created date time
	CreateTime *string `form:"createTime,omitempty" json:"createTime,omitempty" xml:"createTime,omitempty"`
}

// NewAddCouponManagementOK builds a "coupon" service "add" endpoint result
// from a HTTP "OK" response.
func NewAddCouponManagementOK(body *AddResponseBody) *couponviews.CouponManagementView {
	v := &couponviews.CouponManagementView{
		UserID:     body.UserID,
		ID:         body.ID,
		Name:       body.Name,
		Status:     body.Status,
		CreateTime: body.CreateTime,
	}

	return v
}