// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type ListModelsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only models with a creation time greater than or equal
	// to the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only models created before the specified time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of models to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the training job name. This filter returns only models in the
	// training job whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the response to a previous ListModels request was truncated, the response
	// includes a NextToken. To retrieve the next set of models, use the token in
	// the next request.
	NextToken *string `type:"string"`

	// Sorts the list of results. The default is CreationTime.
	SortBy enums.ModelSortKey `type:"string" enum:"true"`

	// The sort order for results. The default is Descending.
	SortOrder enums.OrderKey `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListModelsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListModelsOutput struct {
	_ struct{} `type:"structure"`

	// An array of ModelSummary objects, each of which lists a model.
	//
	// Models is a required field
	Models []ModelSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of models, use it in the subsequent request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListModelsOutput) String() string {
	return awsutil.Prettify(s)
}
