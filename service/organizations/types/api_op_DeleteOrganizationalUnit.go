// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteOrganizationalUnitInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the organizational unit that you want to delete.
	// You can get the ID from the ListOrganizationalUnitsForParent operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational
	// unit ID string requires "ou-" followed by from 4 to 32 lower-case letters
	// or digits (the ID of the root that contains the OU) followed by a second
	// "-" dash and from 8 to 32 additional lower-case letters or digits.
	//
	// OrganizationalUnitId is a required field
	OrganizationalUnitId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOrganizationalUnitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOrganizationalUnitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOrganizationalUnitInput"}

	if s.OrganizationalUnitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationalUnitId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOrganizationalUnitOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationalUnitOutput) String() string {
	return awsutil.Prettify(s)
}