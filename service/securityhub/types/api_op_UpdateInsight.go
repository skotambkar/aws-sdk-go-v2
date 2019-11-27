// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateInsightInput struct {
	_ struct{} `type:"structure"`

	// The updated filters that define this insight.
	Filters *AwsSecurityFindingFilters `type:"structure"`

	// The updated GroupBy attribute that defines this insight.
	GroupByAttribute *string `type:"string"`

	// The ARN of the insight that you want to update.
	//
	// InsightArn is a required field
	InsightArn *string `location:"uri" locationName:"InsightArn" type:"string" required:"true"`

	// The updated name for the insight.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateInsightInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateInsightInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateInsightInput"}

	if s.InsightArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InsightArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateInsightOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateInsightOutput) String() string {
	return awsutil.Prettify(s)
}
