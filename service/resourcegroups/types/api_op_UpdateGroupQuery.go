// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateGroupQueryInput struct {
	_ struct{} `type:"structure"`

	// The name of the resource group for which you want to edit the query.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The resource query that determines which AWS resources are members of the
	// resource group.
	//
	// ResourceQuery is a required field
	ResourceQuery *ResourceQuery `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateGroupQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupQueryInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.ResourceQuery == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceQuery"))
	}
	if s.ResourceQuery != nil {
		if err := s.ResourceQuery.Validate(); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateGroupQueryOutput struct {
	_ struct{} `type:"structure"`

	// The resource query associated with the resource group after the update.
	GroupQuery *GroupQuery `type:"structure"`
}

// String returns the string representation
func (s UpdateGroupQueryOutput) String() string {
	return awsutil.Prettify(s)
}