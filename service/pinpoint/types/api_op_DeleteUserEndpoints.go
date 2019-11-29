// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteUserEndpointsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// UserId is a required field
	UserId *string `location:"uri" locationName:"user-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserEndpointsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserEndpointsOutput struct {
	_ struct{} `type:"structure" payload:"EndpointsResponse"`

	// Provides information about all the endpoints that are associated with a user
	// ID.
	//
	// EndpointsResponse is a required field
	EndpointsResponse *EndpointsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteUserEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}
