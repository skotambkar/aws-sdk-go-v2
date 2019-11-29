// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetJourneyExecutionActivityMetricsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// JourneyActivityId is a required field
	JourneyActivityId *string `location:"uri" locationName:"journey-activity-id" type:"string" required:"true"`

	// JourneyId is a required field
	JourneyId *string `location:"uri" locationName:"journey-id" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`
}

// String returns the string representation
func (s GetJourneyExecutionActivityMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJourneyExecutionActivityMetricsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJourneyExecutionActivityMetricsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.JourneyActivityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JourneyActivityId"))
	}

	if s.JourneyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JourneyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetJourneyExecutionActivityMetricsOutput struct {
	_ struct{} `type:"structure" payload:"JourneyExecutionActivityMetricsResponse"`

	// Provides the results of a query that retrieved the data for a standard execution
	// metric that applies to a journey activity, and provides information about
	// that query.
	//
	// JourneyExecutionActivityMetricsResponse is a required field
	JourneyExecutionActivityMetricsResponse *JourneyExecutionActivityMetricsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetJourneyExecutionActivityMetricsOutput) String() string {
	return awsutil.Prettify(s)
}
