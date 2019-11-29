// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteActionTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the custom action target to delete.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `location:"uri" locationName:"ActionTargetArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActionTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteActionTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteActionTargetInput"}

	if s.ActionTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionTargetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteActionTargetOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the custom action target that was deleted.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActionTargetOutput) String() string {
	return awsutil.Prettify(s)
}
