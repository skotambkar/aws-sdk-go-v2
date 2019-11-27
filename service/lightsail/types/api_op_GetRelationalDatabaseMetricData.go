// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/enums"
)

type GetRelationalDatabaseMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The end of the time interval from which to get metric data.
	//
	// Constraints:
	//
	//    * Specified in Coordinated Universal Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use an
	//    end time of October 1, 2018, at 8 PM UTC, then you input 1538424000 as
	//    the end time.
	//
	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" required:"true"`

	// The name of the metric data to return.
	//
	// MetricName is a required field
	MetricName enums.RelationalDatabaseMetricName `locationName:"metricName" type:"string" required:"true" enum:"true"`

	// The granularity, in seconds, of the returned data points.
	//
	// Period is a required field
	Period *int64 `locationName:"period" min:"60" type:"integer" required:"true"`

	// The name of your database from which to get metric data.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The start of the time interval from which to get metric data.
	//
	// Constraints:
	//
	//    * Specified in Coordinated Universal Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use a
	//    start time of October 1, 2018, at 8 PM UTC, then you input 1538424000
	//    as the start time.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" required:"true"`

	// The array of statistics for your metric data request.
	//
	// Statistics is a required field
	Statistics []enums.MetricStatistic `locationName:"statistics" type:"list" required:"true"`

	// The unit for the metric data request.
	//
	// Unit is a required field
	Unit enums.MetricUnit `locationName:"unit" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseMetricDataInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if len(s.MetricName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Period == nil {
		invalidParams.Add(aws.NewErrParamRequired("Period"))
	}
	if s.Period != nil && *s.Period < 60 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 60))
	}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if s.Statistics == nil {
		invalidParams.Add(aws.NewErrParamRequired("Statistics"))
	}
	if len(s.Unit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Unit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRelationalDatabaseMetricDataOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your get relational database metric data
	// request.
	MetricData []MetricDatapoint `locationName:"metricData" type:"list"`

	// The name of the metric.
	MetricName enums.RelationalDatabaseMetricName `locationName:"metricName" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}
