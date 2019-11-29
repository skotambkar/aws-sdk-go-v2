// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutRemediationConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// A list of remediation configuration objects.
	//
	// RemediationConfigurations is a required field
	RemediationConfigurations []RemediationConfiguration `type:"list" required:"true"`
}

// String returns the string representation
func (s PutRemediationConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRemediationConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRemediationConfigurationsInput"}

	if s.RemediationConfigurations == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemediationConfigurations"))
	}
	if s.RemediationConfigurations != nil {
		for i, v := range s.RemediationConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RemediationConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutRemediationConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// Returns a list of failed remediation batch objects.
	FailedBatches []FailedRemediationBatch `type:"list"`
}

// String returns the string representation
func (s PutRemediationConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}