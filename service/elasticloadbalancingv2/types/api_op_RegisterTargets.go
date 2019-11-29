// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterTargetsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the target group.
	//
	// TargetGroupArn is a required field
	TargetGroupArn *string `type:"string" required:"true"`

	// The targets.
	//
	// To register a target by instance ID, specify the instance ID. To register
	// a target by IP address, specify the IP address. To register a Lambda function,
	// specify the ARN of the Lambda function.
	//
	// Targets is a required field
	Targets []TargetDescription `type:"list" required:"true"`
}

// String returns the string representation
func (s RegisterTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterTargetsInput"}

	if s.TargetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupArn"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
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

type RegisterTargetsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterTargetsOutput) String() string {
	return awsutil.Prettify(s)
}
