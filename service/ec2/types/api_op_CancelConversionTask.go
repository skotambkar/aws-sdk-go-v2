// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CancelConversionTaskInput struct {
	_ struct{} `type:"structure"`

	// The ID of the conversion task.
	//
	// ConversionTaskId is a required field
	ConversionTaskId *string `locationName:"conversionTaskId" type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The reason for canceling the conversion task.
	ReasonMessage *string `locationName:"reasonMessage" type:"string"`
}

// String returns the string representation
func (s CancelConversionTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelConversionTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelConversionTaskInput"}

	if s.ConversionTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConversionTaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelConversionTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelConversionTaskOutput) String() string {
	return awsutil.Prettify(s)
}