// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterDomainInput struct {
	_ struct{} `type:"structure"`

	// A text description of the domain.
	Description *string `locationName:"description" type:"string"`

	// Name of the domain to register. The name must be unique in the region that
	// the domain is registered in.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// Tags to be added when registering a domain.
	//
	// Tags may only contain unicode letters, digits, whitespace, or these symbols:
	// _ . : / = + - @.
	Tags []ResourceTag `locationName:"tags" type:"list"`

	// The duration (in days) that records and histories of workflow executions
	// on the domain should be kept by the service. After the retention period,
	// the workflow execution isn't available in the results of visibility calls.
	//
	// If you pass the value NONE or 0 (zero), then the workflow execution history
	// isn't retained. As soon as the workflow execution completes, the execution
	// record and its history are deleted.
	//
	// The maximum workflow execution retention period is 90 days. For more information
	// about Amazon SWF service limits, see: Amazon SWF Service Limits (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html)
	// in the Amazon SWF Developer Guide.
	//
	// WorkflowExecutionRetentionPeriodInDays is a required field
	WorkflowExecutionRetentionPeriodInDays *string `locationName:"workflowExecutionRetentionPeriodInDays" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterDomainInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.WorkflowExecutionRetentionPeriodInDays == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowExecutionRetentionPeriodInDays"))
	}
	if s.WorkflowExecutionRetentionPeriodInDays != nil && len(*s.WorkflowExecutionRetentionPeriodInDays) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkflowExecutionRetentionPeriodInDays", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterDomainOutput) String() string {
	return awsutil.Prettify(s)
}
