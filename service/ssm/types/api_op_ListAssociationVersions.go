// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAssociationVersionsInput struct {
	_ struct{} `type:"structure"`

	// The association ID for which you want to view all versions.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAssociationVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssociationVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssociationVersionsInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAssociationVersionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about all versions of the association for the specified association
	// ID.
	AssociationVersions []AssociationVersionInfo `min:"1" type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAssociationVersionsOutput) String() string {
	return awsutil.Prettify(s)
}