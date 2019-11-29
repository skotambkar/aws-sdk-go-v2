// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListGroupCertificateAuthoritiesInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListGroupCertificateAuthoritiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupCertificateAuthoritiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupCertificateAuthoritiesInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListGroupCertificateAuthoritiesOutput struct {
	_ struct{} `type:"structure"`

	// A list of certificate authorities associated with the group.
	GroupCertificateAuthorities []GroupCertificateAuthorityProperties `type:"list"`
}

// String returns the string representation
func (s ListGroupCertificateAuthoritiesOutput) String() string {
	return awsutil.Prettify(s)
}
