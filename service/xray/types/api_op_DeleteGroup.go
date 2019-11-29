// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the group that was generated on creation.
	GroupARN *string `min:"1" type:"string"`

	// The case-sensitive name of the group.
	GroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupInput"}
	if s.GroupARN != nil && len(*s.GroupARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupARN", 1))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGroupOutput) String() string {
	return awsutil.Prettify(s)
}