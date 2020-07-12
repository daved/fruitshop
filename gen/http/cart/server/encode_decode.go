// Code generated by goa v3.1.3, DO NOT EDIT.
//
// cart HTTP server encoders and decoders
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	"context"
	cartviews "fruitshop/gen/cart/views"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the cart add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeAddRequest returns a decoder for requests sent to the cart add
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
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			cartID string

			params = mux.Vars(r)
		)
		cartID = params["cartId"]
		payload := NewAddPayload(&body, cartID)

		return payload, nil
	}
}

// EncodeGetResponse returns an encoder for responses returned by the cart get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(cartviews.CartManagementCollection)
		enc := encoder(ctx, w)
		body := NewCartManagementResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the cart get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			cartID string

			params = mux.Vars(r)
		)
		cartID = params["cartId"]
		payload := NewGetPayload(cartID)

		return payload, nil
	}
}

// marshalCartviewsCartManagementViewToCartManagementResponse builds a value of
// type *CartManagementResponse from a value of type
// *cartviews.CartManagementView.
func marshalCartviewsCartManagementViewToCartManagementResponse(v *cartviews.CartManagementView) *CartManagementResponse {
	res := &CartManagementResponse{
		CartID: *v.CartID,
		Name:   *v.Name,
		Count:  *v.Count,
	}

	return res
}