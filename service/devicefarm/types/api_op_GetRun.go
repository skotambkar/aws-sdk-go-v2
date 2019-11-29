// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the get run operation.
type GetRunInput struct {
	_ struct{} `type:"structure"`

	// The run's ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRunInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a get run request.
type GetRunOutput struct {
	_ struct{} `type:"structure"`

	// The run you wish to get results from.
	Run *Run `locationName:"run" type:"structure"`
}

// String returns the string representation
func (s GetRunOutput) String() string {
	return awsutil.Prettify(s)
}
