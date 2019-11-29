// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appstream/enums"
)

type DescribeSessionsInput struct {
	_ struct{} `type:"structure"`

	// The authentication method. Specify API for a user authenticated using a streaming
	// URL or SAML for a SAML federated user. The default is to authenticate users
	// using a streaming URL.
	AuthenticationType enums.AuthenticationType `type:"string" enum:"true"`

	// The name of the fleet. This value is case-sensitive.
	//
	// FleetName is a required field
	FleetName *string `min:"1" type:"string" required:"true"`

	// The size of each page of results. The default value is 20 and the maximum
	// value is 50.
	Limit *int64 `type:"integer"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`

	// The name of the stack. This value is case-sensitive.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The user identifier.
	UserId *string `min:"2" type:"string"`
}

// String returns the string representation
func (s DescribeSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSessionsInput"}

	if s.FleetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetName"))
	}
	if s.FleetName != nil && len(*s.FleetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetName", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSessionsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`

	// Information about the streaming sessions.
	Sessions []Session `type:"list"`
}

// String returns the string representation
func (s DescribeSessionsOutput) String() string {
	return awsutil.Prettify(s)
}