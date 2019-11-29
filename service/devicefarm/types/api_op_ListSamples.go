// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list samples operation.
type ListSamplesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the job used to list samples.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListSamplesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSamplesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSamplesInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list samples request.
type ListSamplesOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned, which can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the samples.
	Samples []Sample `locationName:"samples" type:"list"`
}

// String returns the string representation
func (s ListSamplesOutput) String() string {
	return awsutil.Prettify(s)
}