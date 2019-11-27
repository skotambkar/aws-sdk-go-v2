// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSubscribedWorkteamsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the work team name. This filter returns only work teams whose
	// name contains the specified string.
	NameContains *string `min:"1" type:"string"`

	// If the result of the previous ListSubscribedWorkteams request was truncated,
	// the response includes a NextToken. To retrieve the next set of labeling jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSubscribedWorkteamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSubscribedWorkteamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSubscribedWorkteamsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSubscribedWorkteamsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string `type:"string"`

	// An array of Workteam objects, each describing a work team.
	//
	// SubscribedWorkteams is a required field
	SubscribedWorkteams []SubscribedWorkteam `type:"list" required:"true"`
}

// String returns the string representation
func (s ListSubscribedWorkteamsOutput) String() string {
	return awsutil.Prettify(s)
}
