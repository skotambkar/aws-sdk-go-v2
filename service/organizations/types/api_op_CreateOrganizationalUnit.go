// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateOrganizationalUnitInput struct {
	_ struct{} `type:"structure"`

	// The friendly name to assign to the new OU.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The unique identifier (ID) of the parent root or OU that you want to create
	// the new OU in.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lower-case
	//    letters or digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lower-case letters or digits (the ID of the root that
	//    the OU is in) followed by a second "-" dash and from 8 to 32 additional
	//    lower-case letters or digits.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOrganizationalUnitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOrganizationalUnitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOrganizationalUnitInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateOrganizationalUnitOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the newly created OU.
	OrganizationalUnit *OrganizationalUnit `type:"structure"`
}

// String returns the string representation
func (s CreateOrganizationalUnitOutput) String() string {
	return awsutil.Prettify(s)
}