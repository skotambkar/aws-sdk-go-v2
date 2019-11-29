// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyTargetGroupAttributesInput struct {
	_ struct{} `type:"structure"`

	// The attributes.
	//
	// Attributes is a required field
	Attributes []TargetGroupAttribute `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the target group.
	//
	// TargetGroupArn is a required field
	TargetGroupArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyTargetGroupAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTargetGroupAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyTargetGroupAttributesInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.TargetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyTargetGroupAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the attributes.
	Attributes []TargetGroupAttribute `type:"list"`
}

// String returns the string representation
func (s ModifyTargetGroupAttributesOutput) String() string {
	return awsutil.Prettify(s)
}