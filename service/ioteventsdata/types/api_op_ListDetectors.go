// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDetectorsInput struct {
	_ struct{} `type:"structure"`

	// The name of the detector model whose detectors (instances) are listed.
	//
	// DetectorModelName is a required field
	DetectorModelName *string `location:"uri" locationName:"detectorModelName" min:"1" type:"string" required:"true"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// A filter that limits results to those detectors (instances) in the given
	// state.
	StateName *string `location:"querystring" locationName:"stateName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDetectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDetectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDetectorsInput"}

	if s.DetectorModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorModelName"))
	}
	if s.DetectorModelName != nil && len(*s.DetectorModelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.StateName != nil && len(*s.StateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDetectorsOutput struct {
	_ struct{} `type:"structure"`

	// A list of summary information about the detectors (instances).
	DetectorSummaries []DetectorSummary `locationName:"detectorSummaries" type:"list"`

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorsOutput) String() string {
	return awsutil.Prettify(s)
}