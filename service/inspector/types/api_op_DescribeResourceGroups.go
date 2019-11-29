// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeResourceGroupsInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the resource group that you want to describe.
	//
	// ResourceGroupArns is a required field
	ResourceGroupArns []string `locationName:"resourceGroupArns" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeResourceGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeResourceGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeResourceGroupsInput"}

	if s.ResourceGroupArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupArns"))
	}
	if s.ResourceGroupArns != nil && len(s.ResourceGroupArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeResourceGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Resource group details that cannot be described. An error code is provided
	// for each failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`

	// Information about a resource group.
	//
	// ResourceGroups is a required field
	ResourceGroups []ResourceGroup `locationName:"resourceGroups" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeResourceGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
