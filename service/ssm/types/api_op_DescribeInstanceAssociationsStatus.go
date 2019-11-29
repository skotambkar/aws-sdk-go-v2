// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInstanceAssociationsStatusInput struct {
	_ struct{} `type:"structure"`

	// The instance IDs for which you want association status information.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAssociationsStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceAssociationsStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceAssociationsStatusInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInstanceAssociationsStatusOutput struct {
	_ struct{} `type:"structure"`

	// Status information about the association.
	InstanceAssociationStatusInfos []InstanceAssociationStatusInfo `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAssociationsStatusOutput) String() string {
	return awsutil.Prettify(s)
}
