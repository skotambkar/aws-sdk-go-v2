// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSubscribedWorkteamInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the subscribed work team to describe.
	//
	// WorkteamArn is a required field
	WorkteamArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSubscribedWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSubscribedWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSubscribedWorkteamInput"}

	if s.WorkteamArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSubscribedWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// A Workteam instance that contains information about the work team.
	//
	// SubscribedWorkteam is a required field
	SubscribedWorkteam *SubscribedWorkteam `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeSubscribedWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}