// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRemediationExceptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the AWS Config rule for which you want to delete remediation
	// exception configuration.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	//
	// ResourceKeys is a required field
	ResourceKeys []RemediationExceptionResourceKey `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteRemediationExceptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRemediationExceptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRemediationExceptionsInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if s.ResourceKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceKeys"))
	}
	if s.ResourceKeys != nil && len(s.ResourceKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceKeys", 1))
	}
	if s.ResourceKeys != nil {
		for i, v := range s.ResourceKeys {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceKeys", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRemediationExceptionsOutput struct {
	_ struct{} `type:"structure"`

	// Returns a list of failed delete remediation exceptions batch objects. Each
	// object in the batch consists of a list of failed items and failure messages.
	FailedBatches []FailedDeleteRemediationExceptionsBatch `type:"list"`
}

// String returns the string representation
func (s DeleteRemediationExceptionsOutput) String() string {
	return awsutil.Prettify(s)
}
