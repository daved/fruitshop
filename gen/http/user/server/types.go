// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	user "fruitshop/gen/user"
	userviews "fruitshop/gen/user/views"
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
	// userId
	UserID string `form:"userId" json:"userId" xml:"userId"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// userId
	UserID string `form:"userId" json:"userId" xml:"userId"`
}

// UserManagementResponseCollection is the type of the "user" service "show"
// endpoint HTTP response body.
type UserManagementResponseCollection []*UserManagementResponse

// UserManagementResponse is used to define fields on response body types.
type UserManagementResponse struct {
	// userId
	UserID string `form:"userId" json:"userId" xml:"userId"`
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "user" service.
func NewAddResponseBody(res *userviews.UserManagementView) *AddResponseBody {
	body := &AddResponseBody{
		UserID: *res.UserID,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "user" service.
func NewGetResponseBody(res *userviews.UserManagementView) *GetResponseBody {
	body := &GetResponseBody{
		UserID: *res.UserID,
	}
	return body
}

// NewUserManagementResponseCollection builds the HTTP response body from the
// result of the "show" endpoint of the "user" service.
func NewUserManagementResponseCollection(res userviews.UserManagementCollectionView) UserManagementResponseCollection {
	body := make([]*UserManagementResponse, len(res))
	for i, val := range res {
		body[i] = marshalUserviewsUserManagementViewToUserManagementResponse(val)
	}
	return body
}

// NewAddPayload builds a user service add endpoint payload.
func NewAddPayload(body *AddRequestBody, userID string) *user.AddPayload {
	v := &user.AddPayload{
		ID: body.ID,
	}
	v.UserID = userID

	return v
}

// NewGetPayload builds a user service get endpoint payload.
func NewGetPayload(userID string) *user.GetPayload {
	v := &user.GetPayload{}
	v.UserID = userID

	return v
}
