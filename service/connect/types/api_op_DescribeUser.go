// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUserInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The identifier of the user account.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeUserOutput struct {
	_ struct{} `type:"structure"`

	// Information about the user account and configuration settings.
	User *User `type:"structure"`
}

// String returns the string representation
func (s DescribeUserOutput) String() string {
	return awsutil.Prettify(s)
}