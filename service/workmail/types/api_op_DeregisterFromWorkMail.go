// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterFromWorkMailInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the member (user or group) to be updated.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the Amazon WorkMail entity
	// exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterFromWorkMailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterFromWorkMailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterFromWorkMailInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterFromWorkMailOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterFromWorkMailOutput) String() string {
	return awsutil.Prettify(s)
}
