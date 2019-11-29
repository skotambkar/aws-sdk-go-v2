// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAcceleratorsInput struct {
	_ struct{} `type:"structure"`

	// The number of Global Accelerator objects that you want to return with this
	// call. The default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAcceleratorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAcceleratorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAcceleratorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAcceleratorsOutput struct {
	_ struct{} `type:"structure"`

	// The list of accelerators for a customer account.
	Accelerators []Accelerator `type:"list"`

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAcceleratorsOutput) String() string {
	return awsutil.Prettify(s)
}