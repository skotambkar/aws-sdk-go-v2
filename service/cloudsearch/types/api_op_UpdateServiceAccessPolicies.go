// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the UpdateServiceAccessPolicies operation.
// Specifies the name of the domain you want to update and the access rules
// you want to configure.
type UpdateServiceAccessPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The access rules you want to configure. These rules replace any existing
	// rules.
	//
	// AccessPolicies is a required field
	AccessPolicies *string `type:"string" required:"true"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateServiceAccessPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceAccessPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceAccessPoliciesInput"}

	if s.AccessPolicies == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicies"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of an UpdateServiceAccessPolicies request. Contains the new access
// policies.
type UpdateServiceAccessPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The access rules configured for the domain.
	//
	// AccessPolicies is a required field
	AccessPolicies *AccessPoliciesStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateServiceAccessPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}
