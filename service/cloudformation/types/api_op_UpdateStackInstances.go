// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateStackInstancesInput struct {
	_ struct{} `type:"structure"`

	// The names of one or more AWS accounts for which you want to update parameter
	// values for stack instances. The overridden parameter values will be applied
	// to all stack instances in the specified accounts and regions.
	//
	// Accounts is a required field
	Accounts []string `type:"list" required:"true"`

	// The unique identifier for this stack set operation.
	//
	// The operation ID also functions as an idempotency token, to ensure that AWS
	// CloudFormation performs the stack set operation only once, even if you retry
	// the request multiple times. You might retry stack set operation requests
	// to ensure that AWS CloudFormation successfully received them.
	//
	// If you don't specify an operation ID, the SDK generates one automatically.
	OperationId *string `min:"1" type:"string" idempotencyToken:"true"`

	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *StackSetOperationPreferences `type:"structure"`

	// A list of input parameters whose values you want to update for the specified
	// stack instances.
	//
	// Any overridden parameter values will be applied to all stack instances in
	// the specified accounts and regions. When specifying parameters and their
	// values, be aware of how AWS CloudFormation sets parameter values during stack
	// instance update operations:
	//
	//    * To override the current value for a parameter, include the parameter
	//    and specify its value.
	//
	//    * To leave a parameter set to its present value, you can do one of the
	//    following: Do not include the parameter in the list. Include the parameter
	//    and specify UsePreviousValue as true. (You cannot specify both a value
	//    and set UsePreviousValue to true.)
	//
	//    * To set all overridden parameter back to the values specified in the
	//    stack set, specify a parameter list but do not include any parameters.
	//
	//    * To leave all parameters set to their present values, do not specify
	//    this property at all.
	//
	// During stack set updates, any parameter values overridden for a stack instance
	// are not updated, but retain their overridden value.
	//
	// You can only override the parameter values that are specified in the stack
	// set; to add or delete a parameter itself, use UpdateStackSet to update the
	// stack set template. If you add a parameter to a template, before you can
	// override the parameter value specified in the stack set you must first use
	// UpdateStackSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
	// to update all stack instances with the updated template and parameter value
	// specified in the stack set. Once a stack instance has been updated with the
	// new parameter, you can then override the parameter value using UpdateStackInstances.
	ParameterOverrides []Parameter `type:"list"`

	// The names of one or more regions in which you want to update parameter values
	// for stack instances. The overridden parameter values will be applied to all
	// stack instances in the specified accounts and regions.
	//
	// Regions is a required field
	Regions []string `type:"list" required:"true"`

	// The name or unique ID of the stack set associated with the stack instances.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateStackInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStackInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStackInstancesInput"}

	if s.Accounts == nil {
		invalidParams.Add(aws.NewErrParamRequired("Accounts"))
	}
	if s.OperationId != nil && len(*s.OperationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationId", 1))
	}

	if s.Regions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Regions"))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}
	if s.OperationPreferences != nil {
		if err := s.OperationPreferences.Validate(); err != nil {
			invalidParams.AddNested("OperationPreferences", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateStackInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for this stack set operation.
	OperationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateStackInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
