// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// both objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Reference that identifies the object whose policy object will be detached.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// Reference that identifies the policy object.
	//
	// PolicyReference is a required field
	PolicyReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s DetachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachPolicyInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.PolicyReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
