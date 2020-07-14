// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user HTTP server encoders and decoders
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	"context"
	userviews "fruitshop/gen/user/views"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the user add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*userviews.UserManagement)
		enc := encoder(ctx, w)
		body := NewAddResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the user add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			userID string

			params = mux.Vars(r)
		)
		userID = params["userId"]
		payload := NewAddPayload(&body, userID)

		return payload, nil
	}
}

// EncodeGetResponse returns an encoder for responses returned by the user get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*userviews.UserManagement)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the user get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID string

			params = mux.Vars(r)
		)
		userID = params["userId"]
		payload := NewGetPayload(userID)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the user
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(userviews.UserManagementCollection)
		enc := encoder(ctx, w)
		body := NewUserManagementResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalUserviewsUserManagementViewToUserManagementResponse builds a value of
// type *UserManagementResponse from a value of type
// *userviews.UserManagementView.
func marshalUserviewsUserManagementViewToUserManagementResponse(v *userviews.UserManagementView) *UserManagementResponse {
	res := &UserManagementResponse{
		UserID: *v.UserID,
	}

	return res
}
