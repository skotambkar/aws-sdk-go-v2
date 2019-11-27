// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSampleFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector to create sample findings for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// Types of sample findings that you want to generate.
	FindingTypes []string `locationName:"findingTypes" type:"list"`
}

// String returns the string representation
func (s CreateSampleFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSampleFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSampleFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSampleFindingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSampleFindingsOutput) String() string {
	return awsutil.Prettify(s)
}
