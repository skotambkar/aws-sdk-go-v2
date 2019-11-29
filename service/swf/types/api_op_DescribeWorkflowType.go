// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkflowTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which this workflow type is registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The workflow type to describe.
	//
	// WorkflowType is a required field
	WorkflowType *WorkflowType `locationName:"workflowType" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkflowTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkflowTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkflowTypeInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.WorkflowType == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowType"))
	}
	if s.WorkflowType != nil {
		if err := s.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about a workflow type.
type DescribeWorkflowTypeOutput struct {
	_ struct{} `type:"structure"`

	// Configuration settings of the workflow type registered through RegisterWorkflowType
	//
	// Configuration is a required field
	Configuration *WorkflowTypeConfiguration `locationName:"configuration" type:"structure" required:"true"`

	// General information about the workflow type.
	//
	// The status of the workflow type (returned in the WorkflowTypeInfo structure)
	// can be one of the following.
	//
	//    * REGISTERED – The type is registered and available. Workers supporting
	//    this type should be running.
	//
	//    * DEPRECATED – The type was deprecated using DeprecateWorkflowType,
	//    but is still in use. You should keep workers supporting this type running.
	//    You cannot create new workflow executions of this type.
	//
	// TypeInfo is a required field
	TypeInfo *WorkflowTypeInfo `locationName:"typeInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkflowTypeOutput) String() string {
	return awsutil.Prettify(s)
}