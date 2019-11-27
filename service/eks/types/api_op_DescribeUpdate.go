// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUpdateInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster to update.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The ID of the update to describe.
	//
	// UpdateId is a required field
	UpdateId *string `location:"uri" locationName:"updateId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUpdateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.UpdateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeUpdateOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the specified update.
	Update *Update `locationName:"update" type:"structure"`
}

// String returns the string representation
func (s DescribeUpdateOutput) String() string {
	return awsutil.Prettify(s)
}
