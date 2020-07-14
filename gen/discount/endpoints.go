// Code generated by goa v3.1.3, DO NOT EDIT.
//
// discount endpoints
//
// Command:
// $ goa gen fruitshop/design

package discount

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "discount" service endpoints.
type Endpoints struct {
	Get goa.Endpoint
}

// NewEndpoints wraps the methods of the "discount" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Get: NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "discount" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "discount".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
		res, err := s.Get(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedDiscountManagementCollection(res, "default")
		return vres, nil
	}
}
