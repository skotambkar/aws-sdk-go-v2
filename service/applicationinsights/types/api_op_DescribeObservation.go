// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeObservationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the observation.
	//
	// ObservationId is a required field
	ObservationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeObservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeObservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeObservationInput"}

	if s.ObservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObservationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeObservationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the observation.
	Observation *Observation `type:"structure"`
}

// String returns the string representation
func (s DescribeObservationOutput) String() string {
	return awsutil.Prettify(s)
}