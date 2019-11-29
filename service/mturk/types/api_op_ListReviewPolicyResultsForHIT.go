// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mturk/enums"
)

type ListReviewPolicyResultsForHITInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the HIT to retrieve review results for.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`

	// Limit the number of results returned.
	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination token
	NextToken *string `min:"1" type:"string"`

	// The Policy Level(s) to retrieve review results for - HIT or Assignment. If
	// omitted, the default behavior is to retrieve all data for both policy levels.
	// For a list of all the described policies, see Review Policies.
	PolicyLevels []enums.ReviewPolicyLevel `type:"list"`

	// Specify if the operation should retrieve a list of the actions taken executing
	// the Review Policies and their outcomes.
	RetrieveActions *bool `type:"boolean"`

	// Specify if the operation should retrieve a list of the results computed by
	// the Review Policies.
	RetrieveResults *bool `type:"boolean"`
}

// String returns the string representation
func (s ListReviewPolicyResultsForHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReviewPolicyResultsForHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListReviewPolicyResultsForHITInput"}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListReviewPolicyResultsForHITOutput struct {
	_ struct{} `type:"structure"`

	// The name of the Assignment-level Review Policy. This contains only the PolicyName
	// element.
	AssignmentReviewPolicy *ReviewPolicy `type:"structure"`

	// Contains both ReviewResult and ReviewAction elements for an Assignment.
	AssignmentReviewReport *ReviewReport `type:"structure"`

	// The HITId of the HIT for which results have been returned.
	HITId *string `min:"1" type:"string"`

	// The name of the HIT-level Review Policy. This contains only the PolicyName
	// element.
	HITReviewPolicy *ReviewPolicy `type:"structure"`

	// Contains both ReviewResult and ReviewAction elements for a particular HIT.
	HITReviewReport *ReviewReport `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListReviewPolicyResultsForHITOutput) String() string {
	return awsutil.Prettify(s)
}