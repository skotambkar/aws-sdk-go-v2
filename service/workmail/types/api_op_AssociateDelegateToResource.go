// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDelegateToResourceInput struct {
	_ struct{} `type:"structure"`

	// The member (user or group) to associate to the resource.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The organization under which the resource exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The resource for which members (users or groups) are associated.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDelegateToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDelegateToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDelegateToResourceInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDelegateToResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDelegateToResourceOutput) String() string {
	return awsutil.Prettify(s)
}
