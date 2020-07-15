// Code generated by goa v3.2.0, DO NOT EDIT.
//
// fruit service
//
// Command:
// $ goa gen fruitshop/design

package fruit

import (
	"context"
	fruitviews "fruitshop/gen/fruit/views"
)

// This service allows access to get details about fruits
type Service interface {
	// Get implements get.
	Get(context.Context, *GetPayload) (res *FruitManagement, err error)
	// Show implements show.
	Show(context.Context) (res FruitManagementCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "fruit"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"get", "show"}

// GetPayload is the payload type of the fruit service get method.
type GetPayload struct {
	// Name of the fruit
	Name string
	// Cost of the fruit
	Cost float64
}

// FruitManagement is the result type of the fruit service get method.
type FruitManagement struct {
	// Name is the unique Name of the Fruit.
	Name string
	// Cost of the Fruit.
	Cost float64
}

// FruitManagementCollection is the result type of the fruit service show
// method.
type FruitManagementCollection []*FruitManagement

// NotFound is the type returned when the requested data that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing data
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when the requested data that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return "not_found"
}

// NewFruitManagement initializes result type FruitManagement from viewed
// result type FruitManagement.
func NewFruitManagement(vres *fruitviews.FruitManagement) *FruitManagement {
	return newFruitManagement(vres.Projected)
}

// NewViewedFruitManagement initializes viewed result type FruitManagement from
// result type FruitManagement using the given view.
func NewViewedFruitManagement(res *FruitManagement, view string) *fruitviews.FruitManagement {
	p := newFruitManagementView(res)
	return &fruitviews.FruitManagement{Projected: p, View: "default"}
}

// NewFruitManagementCollection initializes result type
// FruitManagementCollection from viewed result type FruitManagementCollection.
func NewFruitManagementCollection(vres fruitviews.FruitManagementCollection) FruitManagementCollection {
	return newFruitManagementCollection(vres.Projected)
}

// NewViewedFruitManagementCollection initializes viewed result type
// FruitManagementCollection from result type FruitManagementCollection using
// the given view.
func NewViewedFruitManagementCollection(res FruitManagementCollection, view string) fruitviews.FruitManagementCollection {
	p := newFruitManagementCollectionView(res)
	return fruitviews.FruitManagementCollection{Projected: p, View: "default"}
}

// newFruitManagement converts projected type FruitManagement to service type
// FruitManagement.
func newFruitManagement(vres *fruitviews.FruitManagementView) *FruitManagement {
	res := &FruitManagement{}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Cost != nil {
		res.Cost = *vres.Cost
	}
	return res
}

// newFruitManagementView projects result type FruitManagement to projected
// type FruitManagementView using the "default" view.
func newFruitManagementView(res *FruitManagement) *fruitviews.FruitManagementView {
	vres := &fruitviews.FruitManagementView{
		Name: &res.Name,
		Cost: &res.Cost,
	}
	return vres
}

// newFruitManagementCollection converts projected type
// FruitManagementCollection to service type FruitManagementCollection.
func newFruitManagementCollection(vres fruitviews.FruitManagementCollectionView) FruitManagementCollection {
	res := make(FruitManagementCollection, len(vres))
	for i, n := range vres {
		res[i] = newFruitManagement(n)
	}
	return res
}

// newFruitManagementCollectionView projects result type
// FruitManagementCollection to projected type FruitManagementCollectionView
// using the "default" view.
func newFruitManagementCollectionView(res FruitManagementCollection) fruitviews.FruitManagementCollectionView {
	vres := make(fruitviews.FruitManagementCollectionView, len(res))
	for i, n := range res {
		vres[i] = newFruitManagementView(n)
	}
	return vres
}
