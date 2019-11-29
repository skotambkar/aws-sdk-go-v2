// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListPublishingDestinationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector to retrieve publishing destinations for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token to use for paginating results returned in the repsonse. Set the value
	// of this parameter to null for the first request to a list action. For subsequent
	// calls, use the NextToken value returned from the previous request to continue
	// listing results after the first page.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPublishingDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPublishingDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPublishingDestinationsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListPublishingDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// A Destinations obect that includes information about each publishing destination
	// returned.
	//
	// Destinations is a required field
	Destinations []Destination `locationName:"destinations" type:"list" required:"true"`

	// A token to use for paginating results returned in the repsonse. Set the value
	// of this parameter to null for the first request to a list action. For subsequent
	// calls, use the NextToken value returned from the previous request to continue
	// listing results after the first page.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPublishingDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}
