// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutDestinationPolicyInput struct {
	_ struct{} `type:"structure"`

	// An IAM policy document that authorizes cross-account users to deliver their
	// log events to the associated destination.
	//
	// AccessPolicy is a required field
	AccessPolicy *string `locationName:"accessPolicy" min:"1" type:"string" required:"true"`

	// A name for an existing destination.
	//
	// DestinationName is a required field
	DestinationName *string `locationName:"destinationName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutDestinationPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDestinationPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDestinationPolicyInput"}

	if s.AccessPolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicy"))
	}
	if s.AccessPolicy != nil && len(*s.AccessPolicy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessPolicy", 1))
	}

	if s.DestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationName"))
	}
	if s.DestinationName != nil && len(*s.DestinationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutDestinationPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDestinationPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
