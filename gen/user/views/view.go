// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user views
//
// Command:
// $ goa gen fruitshop/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// UserManagement is the viewed result type that is projected based on a view.
type UserManagement struct {
	// Type to project
	Projected *UserManagementView
	// View to render
	View string
}

// UserManagementCollection is the viewed result type that is projected based
// on a view.
type UserManagementCollection struct {
	// Type to project
	Projected UserManagementCollectionView
	// View to render
	View string
}

// UserManagementView is a type that runs validations on a projected type.
type UserManagementView struct {
	// Email ID is the unique id of the User.
	UserEmailID *string
	// Name of the User
	UserName *string
}

// UserManagementCollectionView is a type that runs validations on a projected
// type.
type UserManagementCollectionView []*UserManagementView

var (
	// UserManagementMap is a map of attribute names in result type UserManagement
	// indexed by view name.
	UserManagementMap = map[string][]string{
		"default": []string{
			"UserEmailID",
			"UserName",
		},
	}
	// UserManagementCollectionMap is a map of attribute names in result type
	// UserManagementCollection indexed by view name.
	UserManagementCollectionMap = map[string][]string{
		"default": []string{
			"UserEmailID",
			"UserName",
		},
	}
)

// ValidateUserManagement runs the validations defined on the viewed result
// type UserManagement.
func ValidateUserManagement(result *UserManagement) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserManagementView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateUserManagementCollection runs the validations defined on the viewed
// result type UserManagementCollection.
func ValidateUserManagementCollection(result UserManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateUserManagementView runs the validations defined on
// UserManagementView using the "default" view.
func ValidateUserManagementView(result *UserManagementView) (err error) {
	if result.UserEmailID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UserEmailID", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UserName", "result"))
	}
	return
}

// ValidateUserManagementCollectionView runs the validations defined on
// UserManagementCollectionView using the "default" view.
func ValidateUserManagementCollectionView(result UserManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateUserManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
