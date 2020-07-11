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
	// User Name
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Email ID is the unique id of the User.
	UserEmailID *string `form:"UserEmailID,omitempty" json:"UserEmailID,omitempty" xml:"UserEmailID,omitempty"`
	// Name of the User
	UserName *string `form:"UserName,omitempty" json:"UserName,omitempty" xml:"UserName,omitempty"`
}

// ShowResponseBody is the type of the "user" service "show" endpoint HTTP
// response body.
type ShowResponseBody []*UserManagementResponse

// UserManagementResponse is used to define fields on response body types.
type UserManagementResponse struct {
	// Email ID is the unique id of the User.
	UserEmailID *string `form:"UserEmailID,omitempty" json:"UserEmailID,omitempty" xml:"UserEmailID,omitempty"`
	// Name of the User
	UserName *string `form:"UserName,omitempty" json:"UserName,omitempty" xml:"UserName,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "user" service.
func NewAddRequestBody(p *user.AddPayload) *AddRequestBody {
	body := &AddRequestBody{
		UserName: p.UserName,
	}
	return body
}

// NewGetUserManagementOK builds a "user" service "get" endpoint result from a
// HTTP "OK" response.
func NewGetUserManagementOK(body *GetResponseBody) *userviews.UserManagementView {
	v := &userviews.UserManagementView{
		UserEmailID: body.UserEmailID,
		UserName:    body.UserName,
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
	if body.UserEmailID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UserEmailID", "body"))
	}
	if body.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UserName", "body"))
	}
	return
}
