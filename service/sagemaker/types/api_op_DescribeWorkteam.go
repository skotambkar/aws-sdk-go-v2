// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkteamInput struct {
	_ struct{} `type:"structure"`

	// The name of the work team to return a description of.
	//
	// WorkteamName is a required field
	WorkteamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkteamInput"}

	if s.WorkteamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamName"))
	}
	if s.WorkteamName != nil && len(*s.WorkteamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkteamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// A Workteam instance that contains information about the work team.
	//
	// Workteam is a required field
	Workteam *Workteam `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}
