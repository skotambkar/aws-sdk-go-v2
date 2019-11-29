// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset group that contains the datasets
	// to list.
	DatasetGroupArn *string `locationName:"datasetGroupArn" type:"string"`

	// The maximum number of datasets to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListDatasetImportJobs for getting
	// the next set of dataset import jobs (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatasetsOutput struct {
	_ struct{} `type:"structure"`

	// An array of Dataset objects. Each object provides metadata information.
	Datasets []DatasetSummary `locationName:"datasets" type:"list"`

	// A token for getting the next set of datasets (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetsOutput) String() string {
	return awsutil.Prettify(s)
}