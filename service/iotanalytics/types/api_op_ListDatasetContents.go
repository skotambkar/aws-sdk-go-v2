// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetContentsInput struct {
	_ struct{} `type:"structure"`

	// The name of the data set whose contents information you want to list.
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"datasetName" min:"1" type:"string" required:"true"`

	// The maximum number of results to return in this request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// A filter to limit results to those data set contents whose creation is scheduled
	// before the given time. See the field triggers.schedule in the CreateDataset
	// request. (timestamp)
	ScheduledBefore *time.Time `location:"querystring" locationName:"scheduledBefore" type:"timestamp"`

	// A filter to limit results to those data set contents whose creation is scheduled
	// on or after the given time. See the field triggers.schedule in the CreateDataset
	// request. (timestamp)
	ScheduledOnOrAfter *time.Time `location:"querystring" locationName:"scheduledOnOrAfter" type:"timestamp"`
}

// String returns the string representation
func (s ListDatasetContentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetContentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetContentsInput"}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatasetContentsOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about data set contents that have been created.
	DatasetContentSummaries []DatasetContentSummary `locationName:"datasetContentSummaries" type:"list"`

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetContentsOutput) String() string {
	return awsutil.Prettify(s)
}