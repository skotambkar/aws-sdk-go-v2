// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type ListOTAUpdatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token used to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The OTA update job status.
	OtaUpdateStatus enums.OTAUpdateStatus `location:"querystring" locationName:"otaUpdateStatus" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListOTAUpdatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOTAUpdatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOTAUpdatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOTAUpdatesOutput struct {
	_ struct{} `type:"structure"`

	// A token to use to get the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of OTA update jobs.
	OtaUpdates []OTAUpdateSummary `locationName:"otaUpdates" type:"list"`
}

// String returns the string representation
func (s ListOTAUpdatesOutput) String() string {
	return awsutil.Prettify(s)
}