// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type StartAutomationExecutionInput struct {
	_ struct{} `type:"structure"`

	// User-provided idempotency token. The token must be unique, is case insensitive,
	// enforces the UUID format, and can't be reused.
	ClientToken *string `min:"36" type:"string"`

	// The name of the Automation document to use for this execution.
	//
	// DocumentName is a required field
	DocumentName *string `type:"string" required:"true"`

	// The version of the Automation document to use for this execution.
	DocumentVersion *string `type:"string"`

	// The maximum number of targets allowed to run this task in parallel. You can
	// specify a number, such as 10, or a percentage, such as 10%. The default value
	// is 10.
	MaxConcurrency *string `min:"1" type:"string"`

	// The number of errors that are allowed before the system stops running the
	// automation on additional targets. You can specify either an absolute number
	// of errors, for example 10, or a percentage of the target set, for example
	// 10%. If you specify 3, for example, the system stops running the automation
	// when the fourth error is received. If you specify 0, then the system stops
	// running the automation on additional targets after the first error result
	// is returned. If you run an automation on 50 resources and set max-errors
	// to 10%, then the system stops running the automation on additional targets
	// when the sixth error is received.
	//
	// Executions that are already running an automation when max-errors is reached
	// are allowed to complete, but some of these executions may fail as well. If
	// you need to ensure that there won't be more than max-errors failed executions,
	// set max-concurrency to 1 so the executions proceed one at a time.
	MaxErrors *string `min:"1" type:"string"`

	// The execution mode of the automation. Valid modes include the following:
	// Auto and Interactive. The default mode is Auto.
	Mode enums.ExecutionMode `type:"string" enum:"true"`

	// A key-value map of execution parameters, which match the declared parameters
	// in the Automation document.
	Parameters map[string][]string `min:"1" type:"map"`

	// A location is a combination of AWS Regions and/or AWS accounts where you
	// want to run the Automation. Use this action to start an Automation in multiple
	// Regions and multiple accounts. For more information, see Executing Automations
	// in Multiple AWS Regions and Accounts (http://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-automation-multiple-accounts-and-regions.html)
	// in the AWS Systems Manager User Guide.
	TargetLocations []TargetLocation `min:"1" type:"list"`

	// A key-value mapping of document parameters to target resources. Both Targets
	// and TargetMaps cannot be specified together.
	TargetMaps []map[string][]string `type:"list"`

	// The name of the parameter used as the target resource for the rate-controlled
	// execution. Required if you specify targets.
	TargetParameterName *string `min:"1" type:"string"`

	// A key-value mapping to target resources. Required if you specify TargetParameterName.
	Targets []Target `type:"list"`
}

// String returns the string representation
func (s StartAutomationExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAutomationExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAutomationExecutionInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if s.DocumentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentName"))
	}
	if s.MaxConcurrency != nil && len(*s.MaxConcurrency) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxConcurrency", 1))
	}
	if s.MaxErrors != nil && len(*s.MaxErrors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxErrors", 1))
	}
	if s.Parameters != nil && len(s.Parameters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Parameters", 1))
	}
	if s.TargetLocations != nil && len(s.TargetLocations) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetLocations", 1))
	}
	if s.TargetParameterName != nil && len(*s.TargetParameterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetParameterName", 1))
	}
	if s.TargetLocations != nil {
		for i, v := range s.TargetLocations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetLocations", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartAutomationExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a newly scheduled automation execution.
	AutomationExecutionId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s StartAutomationExecutionOutput) String() string {
	return awsutil.Prettify(s)
}
