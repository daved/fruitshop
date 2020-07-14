// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	user "fruitshop/gen/user"
	userviews "fruitshop/gen/user/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "user" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// ID
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
}

// AddResponseBody is the type of the "user" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// ID is the unique id of the User.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// userId
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID is the unique id of the User.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// userId
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
}

// ShowResponseBody is the type of the "user" service "show" endpoint HTTP
// response body.
type ShowResponseBody []*UserManagementResponse

// UserManagementResponse is used to define fields on response body types.
type UserManagementResponse struct {
	// ID is the unique id of the User.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// userId
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "user" service.
func NewAddRequestBody(p *user.AddPayload) *AddRequestBody {
	body := &AddRequestBody{
		ID: p.ID,
	}
	return body
}

// NewAddUserManagementCreated builds a "user" service "add" endpoint result
// from a HTTP "Created" response.
func NewAddUserManagementCreated(body *AddResponseBody) *userviews.UserManagementView {
	v := &userviews.UserManagementView{
		ID:     body.ID,
		UserID: body.UserID,
	}

	return v
}

// NewGetUserManagementOK builds a "user" service "get" endpoint result from a
// HTTP "OK" response.
func NewGetUserManagementOK(body *GetResponseBody) *userviews.UserManagementView {
	v := &userviews.UserManagementView{
		ID:     body.ID,
		UserID: body.UserID,
	}

	return v
}

// NewShowUserManagementCollectionOK builds a "user" service "show" endpoint
// result from a HTTP "OK" response.
func NewShowUserManagementCollectionOK(body ShowResponseBody) userviews.UserManagementCollectionView {
	v := make([]*userviews.UserManagementView, len(body))
	for i, val := range body {
		v[i] = unmarshalUserManagementResponseToUserviewsUserManagementView(val)
	}
	return v
}

// ValidateUserManagementResponse runs the validations defined on
// UserManagementResponse
func ValidateUserManagementResponse(body *UserManagementResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ID", "body"))
	}
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userId", "body"))
	}
	return
}
