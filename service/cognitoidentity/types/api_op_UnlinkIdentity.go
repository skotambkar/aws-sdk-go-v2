// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the UnlinkIdentity action.
type UnlinkIdentityInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	//
	// IdentityId is a required field
	IdentityId *string `min:"1" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	//
	// Logins is a required field
	Logins map[string]string `type:"map" required:"true"`

	// Provider names to unlink from this identity.
	//
	// LoginsToRemove is a required field
	LoginsToRemove []string `type:"list" required:"true"`
}

// String returns the string representation
func (s UnlinkIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnlinkIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnlinkIdentityInput"}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.Logins == nil {
		invalidParams.Add(aws.NewErrParamRequired("Logins"))
	}

	if s.LoginsToRemove == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoginsToRemove"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UnlinkIdentityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UnlinkIdentityOutput) String() string {
	return awsutil.Prettify(s)
}
