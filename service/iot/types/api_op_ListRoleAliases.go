// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRoleAliasesInput struct {
	_ struct{} `type:"structure"`

	// Return the list of role aliases in ascending alphabetical order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// A marker used to get the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The maximum number of results to return at one time.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListRoleAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRoleAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRoleAliasesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRoleAliasesOutput struct {
	_ struct{} `type:"structure"`

	// A marker used to get the next set of results.
	NextMarker *string `locationName:"nextMarker" type:"string"`

	// The role aliases.
	RoleAliases []string `locationName:"roleAliases" type:"list"`
}

// String returns the string representation
func (s ListRoleAliasesOutput) String() string {
	return awsutil.Prettify(s)
}