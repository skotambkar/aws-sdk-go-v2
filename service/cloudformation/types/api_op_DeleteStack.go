// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for DeleteStack action.
type DeleteStackInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for this DeleteStack request. Specify this token if you
	// plan to retry requests so that AWS CloudFormation knows that you're not attempting
	// to delete a stack with the same name. You might retry DeleteStack requests
	// to ensure that AWS CloudFormation successfully received them.
	//
	// All events triggered by a given stack operation are assigned the same client
	// request token, which you can use to track operations. For example, if you
	// execute a CreateStack operation with the token token1, then all the StackEvents
	// generated by that operation will have ClientRequestToken set as token1.
	//
	// In the console, stack operations display the client request token on the
	// Events tab. Stack operations that are initiated from the console use the
	// token format Console-StackOperation-ID, which helps you easily identify the
	// stack operation . For example, if you create a stack using the console, each
	// stack event would be assigned the same token in the following format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002.
	ClientRequestToken *string `min:"1" type:"string"`

	// For stacks in the DELETE_FAILED state, a list of resource logical IDs that
	// are associated with the resources you want to retain. During deletion, AWS
	// CloudFormation deletes the stack but does not delete the retained resources.
	//
	// Retaining resources is useful when you cannot delete a resource, such as
	// a non-empty S3 bucket, but you want to delete the stack.
	RetainResources []string `type:"list"`

	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes to delete the stack. AWS CloudFormation
	// uses the role's credentials to make calls on your behalf.
	//
	// If you don't specify a value, AWS CloudFormation uses the role that was previously
	// associated with the stack. If no role is available, AWS CloudFormation uses
	// a temporary session that is generated from your user credentials.
	RoleARN *string `min:"20" type:"string"`

	// The name or the unique stack ID that is associated with the stack.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStackInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 20))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteStackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStackOutput) String() string {
	return awsutil.Prettify(s)
}
