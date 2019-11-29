// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRegexMatchSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of RegexMatchSet objects that you want AWS WAF to return
	// for this request. If you have more RegexMatchSet objects than the number
	// you specify for Limit, the response includes a NextMarker value that you
	// can use to get another batch of RegexMatchSet objects.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more RegexMatchSet objects
	// than the value of Limit, AWS WAF returns a NextMarker value in the response
	// that allows you to list another group of ByteMatchSets. For the second and
	// subsequent ListRegexMatchSets requests, specify the value of NextMarker from
	// the previous response to get information about another batch of RegexMatchSet
	// objects.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRegexMatchSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRegexMatchSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRegexMatchSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRegexMatchSetsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more RegexMatchSet objects than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more RegexMatchSet objects, submit another ListRegexMatchSets request, and
	// specify the NextMarker value from the response in the NextMarker value in
	// the next request.
	NextMarker *string `min:"1" type:"string"`

	// An array of RegexMatchSetSummary objects.
	RegexMatchSets []RegexMatchSetSummary `type:"list"`
}

// String returns the string representation
func (s ListRegexMatchSetsOutput) String() string {
	return awsutil.Prettify(s)
}