// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetParametersByPathInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`

	// Filters to limit the request results.
	ParameterFilters []ParameterStringFilter `type:"list"`

	// The hierarchy for the parameter. Hierarchies start with a forward slash (/)
	// and end with the parameter name. A parameter name hierarchy can have a maximum
	// of 15 levels. Here is an example of a hierarchy: /Finance/Prod/IAD/WinServ2016/license33
	//
	// Path is a required field
	Path *string `min:"1" type:"string" required:"true"`

	// Retrieve all parameters within a hierarchy.
	//
	// If a user has access to a path, then the user can access all levels of that
	// path. For example, if a user has permission to access path /a, then the user
	// can also access /a/b. Even if a user has explicitly been denied access in
	// IAM for parameter /a/b, they can still call the GetParametersByPath API action
	// recursively for /a and view /a/b.
	Recursive *bool `type:"boolean"`

	// Retrieve all parameters in a hierarchy with their value decrypted.
	WithDecryption *bool `type:"boolean"`
}

// String returns the string representation
func (s GetParametersByPathInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetParametersByPathInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetParametersByPathInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("Path"))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}
	if s.ParameterFilters != nil {
		for i, v := range s.ParameterFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetParametersByPathOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`

	// A list of parameters found in the specified hierarchy.
	Parameters []Parameter `type:"list"`
}

// String returns the string representation
func (s GetParametersByPathOutput) String() string {
	return awsutil.Prettify(s)
}