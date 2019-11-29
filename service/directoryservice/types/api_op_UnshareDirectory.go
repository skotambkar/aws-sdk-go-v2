// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UnshareDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS Managed Microsoft AD directory that you want to
	// stop sharing.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// Identifier for the directory consumer account with whom the directory has
	// to be unshared.
	//
	// UnshareTarget is a required field
	UnshareTarget *UnshareTarget `type:"structure" required:"true"`
}

// String returns the string representation
func (s UnshareDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnshareDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnshareDirectoryInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.UnshareTarget == nil {
		invalidParams.Add(aws.NewErrParamRequired("UnshareTarget"))
	}
	if s.UnshareTarget != nil {
		if err := s.UnshareTarget.Validate(); err != nil {
			invalidParams.AddNested("UnshareTarget", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UnshareDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// Identifier of the directory stored in the directory consumer account that
	// is to be unshared from the specified directory (DirectoryId).
	SharedDirectoryId *string `type:"string"`
}

// String returns the string representation
func (s UnshareDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}