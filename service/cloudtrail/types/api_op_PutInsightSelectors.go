// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutInsightSelectorsInput struct {
	_ struct{} `type:"structure"`

	// A JSON string that contains the insight types you want to log on a trail.
	// In this release, only ApiCallRateInsight is supported as an insight type.
	//
	// InsightSelectors is a required field
	InsightSelectors []InsightSelector `type:"list" required:"true"`

	// The name of the CloudTrail trail for which you want to change or add Insights
	// selectors.
	//
	// TrailName is a required field
	TrailName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutInsightSelectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutInsightSelectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutInsightSelectorsInput"}

	if s.InsightSelectors == nil {
		invalidParams.Add(aws.NewErrParamRequired("InsightSelectors"))
	}

	if s.TrailName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrailName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutInsightSelectorsOutput struct {
	_ struct{} `type:"structure"`

	// A JSON string that contains the insight types you want to log on a trail.
	// In this release, only ApiCallRateInsight is supported as an insight type.
	InsightSelectors []InsightSelector `type:"list"`

	// The Amazon Resource Name (ARN) of a trail for which you want to change or
	// add Insights selectors.
	TrailARN *string `type:"string"`
}

// String returns the string representation
func (s PutInsightSelectorsOutput) String() string {
	return awsutil.Prettify(s)
}
