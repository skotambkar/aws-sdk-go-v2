// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/enums"
)

// Represents the input of a ListActionTypes action.
type ListActionTypesInput struct {
	_ struct{} `type:"structure"`

	// Filters the list of action types to those created by a specified entity.
	ActionOwnerFilter enums.ActionOwner `locationName:"actionOwnerFilter" type:"string" enum:"true"`

	// An identifier that was returned from the previous list action types call,
	// which can be used to return the next set of action types in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActionTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActionTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActionTypesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListActionTypes action.
type ListActionTypesOutput struct {
	_ struct{} `type:"structure"`

	// Provides details of the action types.
	//
	// ActionTypes is a required field
	ActionTypes []ActionType `locationName:"actionTypes" type:"list" required:"true"`

	// If the amount of returned information is significantly large, an identifier
	// is also returned. It can be used in a subsequent list action types call to
	// return the next set of action types in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActionTypesOutput) String() string {
	return awsutil.Prettify(s)
}
