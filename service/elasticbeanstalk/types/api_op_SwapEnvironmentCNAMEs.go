// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Swaps the CNAMEs of two environments.
type SwapEnvironmentCNAMEsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the destination environment.
	//
	// Condition: You must specify at least the DestinationEnvironmentID or the
	// DestinationEnvironmentName. You may also specify both. You must specify the
	// SourceEnvironmentId with the DestinationEnvironmentId.
	DestinationEnvironmentId *string `type:"string"`

	// The name of the destination environment.
	//
	// Condition: You must specify at least the DestinationEnvironmentID or the
	// DestinationEnvironmentName. You may also specify both. You must specify the
	// SourceEnvironmentName with the DestinationEnvironmentName.
	DestinationEnvironmentName *string `min:"4" type:"string"`

	// The ID of the source environment.
	//
	// Condition: You must specify at least the SourceEnvironmentID or the SourceEnvironmentName.
	// You may also specify both. If you specify the SourceEnvironmentId, you must
	// specify the DestinationEnvironmentId.
	SourceEnvironmentId *string `type:"string"`

	// The name of the source environment.
	//
	// Condition: You must specify at least the SourceEnvironmentID or the SourceEnvironmentName.
	// You may also specify both. If you specify the SourceEnvironmentName, you
	// must specify the DestinationEnvironmentName.
	SourceEnvironmentName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s SwapEnvironmentCNAMEsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SwapEnvironmentCNAMEsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SwapEnvironmentCNAMEsInput"}
	if s.DestinationEnvironmentName != nil && len(*s.DestinationEnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationEnvironmentName", 4))
	}
	if s.SourceEnvironmentName != nil && len(*s.SourceEnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceEnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SwapEnvironmentCNAMEsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SwapEnvironmentCNAMEsOutput) String() string {
	return awsutil.Prettify(s)
}