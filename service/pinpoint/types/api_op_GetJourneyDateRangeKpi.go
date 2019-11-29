// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetJourneyDateRangeKpiInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	EndTime *time.Time `location:"querystring" locationName:"end-time" type:"timestamp" timestampFormat:"iso8601"`

	// JourneyId is a required field
	JourneyId *string `location:"uri" locationName:"journey-id" type:"string" required:"true"`

	// KpiName is a required field
	KpiName *string `location:"uri" locationName:"kpi-name" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	StartTime *time.Time `location:"querystring" locationName:"start-time" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetJourneyDateRangeKpiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJourneyDateRangeKpiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJourneyDateRangeKpiInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.JourneyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JourneyId"))
	}

	if s.KpiName == nil {
		invalidParams.Add(aws.NewErrParamRequired("KpiName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetJourneyDateRangeKpiOutput struct {
	_ struct{} `type:"structure" payload:"JourneyDateRangeKpiResponse"`

	// Provides the results of a query that retrieved the data for a standard engagement
	// metric that applies to a journey, and provides information about that query.
	//
	// JourneyDateRangeKpiResponse is a required field
	JourneyDateRangeKpiResponse *JourneyDateRangeKpiResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetJourneyDateRangeKpiOutput) String() string {
	return awsutil.Prettify(s)
}