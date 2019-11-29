// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrganizationInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the organization to be described.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOrganizationInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// The alias for an organization.
	Alias *string `min:"1" type:"string"`

	// The date at which the organization became usable in the WorkMail context,
	// in UNIX epoch time format.
	CompletedDate *time.Time `type:"timestamp"`

	// The default mail domain associated with the organization.
	DefaultMailDomain *string `type:"string"`

	// The identifier for the directory associated with an Amazon WorkMail organization.
	DirectoryId *string `type:"string"`

	// The type of directory associated with the WorkMail organization.
	DirectoryType *string `type:"string"`

	// (Optional) The error message indicating if unexpected behavior was encountered
	// with regards to the organization.
	ErrorMessage *string `type:"string"`

	// The identifier of an organization.
	OrganizationId *string `type:"string"`

	// The state of an organization.
	State *string `type:"string"`
}

// String returns the string representation
func (s DescribeOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}