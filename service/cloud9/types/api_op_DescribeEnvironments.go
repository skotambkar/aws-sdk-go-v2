// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of individual environments to get information about.
	//
	// EnvironmentIds is a required field
	EnvironmentIds []string `locationName:"environmentIds" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentsInput"}

	if s.EnvironmentIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentIds"))
	}
	if s.EnvironmentIds != nil && len(s.EnvironmentIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the environments that are returned.
	Environments []Environment `locationName:"environments" type:"list"`
}

// String returns the string representation
func (s DescribeEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}