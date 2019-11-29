// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRetirableGrantsInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter to specify the maximum number of items to return. When
	// this value is present, AWS KMS does not return more than the specified number
	// of items, but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and
	// 100, inclusive. If you do not include a value, it defaults to 50.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string `min:"1" type:"string"`

	// The retiring principal for which to list grants.
	//
	// To specify the retiring principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an AWS principal. Valid AWS principals include AWS accounts (root), IAM
	// users, federated users, and assumed role users. For examples of the ARN syntax
	// for specifying a principal, see AWS Identity and Access Management (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the Amazon Web Services General Reference.
	//
	// RetiringPrincipal is a required field
	RetiringPrincipal *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListRetirableGrantsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRetirableGrantsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRetirableGrantsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if s.RetiringPrincipal == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetiringPrincipal"))
	}
	if s.RetiringPrincipal != nil && len(*s.RetiringPrincipal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RetiringPrincipal", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRetirableGrantsOutput struct {
	_ struct{} `type:"structure"`

	// A list of grants.
	Grants []GrantListEntry `type:"list"`

	// When Truncated is true, this element is present and contains the value to
	// use for the Marker parameter in a subsequent request.
	NextMarker *string `min:"1" type:"string"`

	// A flag that indicates whether there are more items in the list. When this
	// value is true, the list in this response is truncated. To get more items,
	// pass the value of the NextMarker element in thisresponse to the Marker parameter
	// in a subsequent request.
	Truncated *bool `type:"boolean"`
}

// String returns the string representation
func (s ListRetirableGrantsOutput) String() string {
	return awsutil.Prettify(s)
}