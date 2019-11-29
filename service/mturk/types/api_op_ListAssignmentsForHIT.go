// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mturk/enums"
)

type ListAssignmentsForHITInput struct {
	_ struct{} `type:"structure"`

	// The status of the assignments to return: Submitted | Approved | Rejected
	AssignmentStatuses []enums.AssignmentStatus `type:"list"`

	// The ID of the HIT.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`

	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination token
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssignmentsForHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssignmentsForHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssignmentsForHITInput"}

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

type ListAssignmentsForHITOutput struct {
	_ struct{} `type:"structure"`

	// The collection of Assignment data structures returned by this call.
	Assignments []Assignment `type:"list"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int64 `type:"integer"`
}

// String returns the string representation
func (s ListAssignmentsForHITOutput) String() string {
	return awsutil.Prettify(s)
}
