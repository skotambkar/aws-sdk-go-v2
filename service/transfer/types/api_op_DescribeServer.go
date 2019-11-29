// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeServerInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServerInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeServerOutput struct {
	_ struct{} `type:"structure"`

	// An array containing the properties of the server with the ServerID you specified.
	//
	// Server is a required field
	Server *DescribedServer `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeServerOutput) String() string {
	return awsutil.Prettify(s)
}