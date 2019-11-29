// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListStackResource action.
type ListStackResourcesInput struct {
	_ struct{} `type:"structure"`

	// A string that identifies the next page of stack resources that you want to
	// retrieve.
	NextToken *string `min:"1" type:"string"`

	// The name or the unique stack ID that is associated with the stack, which
	// are not always interchangeable:
	//
	//    * Running stacks: You can specify either the stack's name or its unique
	//    stack ID.
	//
	//    * Deleted stacks: You must specify the unique stack ID.
	//
	// Default: There is no default value.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListStackResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackResourcesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a ListStackResources action.
type ListStackResourcesOutput struct {
	_ struct{} `type:"structure"`

	// If the output exceeds 1 MB, a string that identifies the next page of stack
	// resources. If no additional page exists, this value is null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackResourceSummary structures.
	StackResourceSummaries []StackResourceSummary `type:"list"`
}

// String returns the string representation
func (s ListStackResourcesOutput) String() string {
	return awsutil.Prettify(s)
}