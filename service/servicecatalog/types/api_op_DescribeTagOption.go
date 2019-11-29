// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTagOptionInput struct {
	_ struct{} `type:"structure"`

	// The TagOption identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTagOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTagOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTagOptionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTagOptionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the TagOption.
	TagOptionDetail *TagOptionDetail `type:"structure"`
}

// String returns the string representation
func (s DescribeTagOptionOutput) String() string {
	return awsutil.Prettify(s)
}