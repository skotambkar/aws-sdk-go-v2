// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRouteInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// RouteName is a required field
	RouteName *string `location:"uri" locationName:"routeName" min:"1" type:"string" required:"true"`

	// An object representing the specification of a route.
	//
	// Spec is a required field
	Spec *RouteSpec `locationName:"spec" type:"structure" required:"true"`

	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRouteInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.RouteName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteName"))
	}
	if s.RouteName != nil && len(*s.RouteName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RouteName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRouteOutput struct {
	_ struct{} `type:"structure" payload:"Route"`

	// An object representing a route returned by a describe operation.
	//
	// Route is a required field
	Route *RouteData `locationName:"route" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateRouteOutput) String() string {
	return awsutil.Prettify(s)
}
