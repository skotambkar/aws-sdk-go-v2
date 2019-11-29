// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartRemediationExecutionInput struct {
	_ struct{} `type:"structure"`

	// The list of names of AWS Config rules that you want to run remediation execution
	// for.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// ResourceKeys is a required field
	ResourceKeys []ResourceKey `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s StartRemediationExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartRemediationExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartRemediationExecutionInput"}

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

type StartRemediationExecutionOutput struct {
	_ struct{} `type:"structure"`

	// For resources that have failed to start execution, the API returns a resource
	// key object.
	FailedItems []ResourceKey `min:"1" type:"list"`

	// Returns a failure message. For example, the resource is already compliant.
	FailureMessage *string `type:"string"`
}

// String returns the string representation
func (s StartRemediationExecutionOutput) String() string {
	return awsutil.Prettify(s)
}