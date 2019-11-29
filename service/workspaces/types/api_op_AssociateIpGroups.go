// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateIpGroupsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory.
	//
	// DirectoryId is a required field
	DirectoryId *string `min:"10" type:"string" required:"true"`

	// The identifiers of one or more IP access control groups.
	//
	// GroupIds is a required field
	GroupIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateIpGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateIpGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateIpGroupsInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.DirectoryId != nil && len(*s.DirectoryId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("DirectoryId", 10))
	}

	if s.GroupIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateIpGroupsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateIpGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
