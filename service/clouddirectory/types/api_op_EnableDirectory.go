// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the directory to enable.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableDirectoryInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the enabled directory.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}