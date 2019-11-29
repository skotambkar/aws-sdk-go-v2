// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListFragmentsInput struct {
	_ struct{} `type:"structure"`

	// Describes the timestamp range and timestamp origin for the range of fragments
	// to return.
	FragmentSelector *FragmentSelector `type:"structure"`

	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a ListFragmentsOutput$NextToken
	// is provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"long"`

	// A token to specify where to start paginating. This is the ListFragmentsOutput$NextToken
	// from a previously truncated response.
	NextToken *string `min:"1" type:"string"`

	// The name of the stream from which to retrieve a fragment list.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFragmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFragmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFragmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}
	if s.FragmentSelector != nil {
		if err := s.FragmentSelector.Validate(); err != nil {
			invalidParams.AddNested("FragmentSelector", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListFragmentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of archived Fragment objects from the stream that meet the selector
	// criteria. Results are in no specific order, even across pages.
	Fragments []Fragment `type:"list"`

	// If the returned list is truncated, the operation returns this token to use
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFragmentsOutput) String() string {
	return awsutil.Prettify(s)
}