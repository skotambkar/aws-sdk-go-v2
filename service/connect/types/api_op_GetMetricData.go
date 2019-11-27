// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/connect/enums"
)

type GetMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The timestamp, in UNIX Epoch time format, at which to end the reporting interval
	// for the retrieval of historical metrics data. The time must be specified
	// using an interval of 5 minutes, such as 11:00, 11:05, 11:10, and must be
	// later than the StartTime timestamp.
	//
	// The time range between StartTime and EndTime must be less than 24 hours.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// A Filters object that contains a list of queue IDs or queue ARNs, up to 100,
	// or a list of Channels to use to filter the metrics returned in the response.
	// Metric data is retrieved only for the resources associated with the IDs,
	// ARNs, or Channels included in the filter. You can use both IDs and ARNs together
	// in a request. Only VOICE is supported for Channel.
	//
	// To find the ARN for a queue, open the queue you want to use in the Amazon
	// Connect Queue editor. The ARN for the queue is displayed in the address bar
	// as part of the URL. For example, the queue ARN is the set of characters at
	// the end of the URL, after 'id=' such as arn:aws:connect:us-east-1:270923740243:instance/78fb859d-1b7d-44b1-8aa3-12f0835c5855/queue/1d1a4575-9618-40ab-bbeb-81e45795fe61.
	// The queue ID is also included in the URL, and is the string after 'queue/'.
	//
	// Filters is a required field
	Filters *Filters `type:"structure" required:"true"`

	// The grouping applied to the metrics returned. For example, when results are
	// grouped by queueId, the metrics returned are grouped by queue. The values
	// returned apply to the metrics for each queue rather than aggregated for all
	// queues.
	//
	// The current version supports grouping by Queue
	//
	// If no Grouping is included in the request, a summary of HistoricalMetrics
	// for all queues is returned.
	Groupings []enums.Grouping `type:"list"`

	// A list of HistoricalMetric objects that contain the metrics to retrieve with
	// the request.
	//
	// A HistoricalMetric object contains: HistoricalMetricName, Statistic, Threshold,
	// and Unit.
	//
	// You must list each metric to retrieve data for in the request. For each historical
	// metric you include in the request, you must include a Unit and a Statistic.
	//
	// The following historical metrics are available:
	//
	// CONTACTS_QUEUED
	//
	// Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_HANDLED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_ABANDONED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_CONSULTED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_AGENT_HUNG_UP_FIRST
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_HANDLED_INCOMING
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_HANDLED_OUTBOUND
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_HOLD_ABANDONS
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_TRANSFERRED_IN
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_TRANSFERRED_OUT
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_TRANSFERRED_IN_FROM_QUEUE
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_TRANSFERRED_OUT_FROM_QUEUE
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CALLBACK_CONTACTS_HANDLED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CALLBACK_CONTACTS_HANDLED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// API_CONTACTS_HANDLED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// CONTACTS_MISSED
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// OCCUPANCY
	//
	// Unit: PERCENT
	//
	// Statistics: AVG
	//
	// HANDLE_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// AFTER_CONTACT_WORK_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// QUEUED_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: MAX
	//
	// ABANDON_TIME
	//
	// Unit: COUNT
	//
	// Statistics: SUM
	//
	// QUEUE_ANSWER_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// HOLD_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// INTERACTION_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// INTERACTION_AND_HOLD_TIME
	//
	// Unit: SECONDS
	//
	// Statistics: AVG
	//
	// SERVICE_LEVEL
	//
	// Unit: PERCENT
	//
	// Statistics: AVG
	//
	// Threshold: Only "Less than" comparisons are supported, with the following
	// service level thresholds: 15, 20, 25, 30, 45, 60, 90, 120, 180, 240, 300,
	// 600
	//
	// HistoricalMetrics is a required field
	HistoricalMetrics []HistoricalMetric `type:"list" required:"true"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// Indicates the maximum number of results to return per page in the response,
	// between 1-100.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `type:"string"`

	// The timestamp, in UNIX Epoch time format, at which to start the reporting
	// interval for the retrieval of historical metrics data. The time must be specified
	// using a multiple of 5 minutes, such as 10:05, 10:10, 10:15.
	//
	// StartTime cannot be earlier than 24 hours before the time of the request.
	// Historical metrics are available in Amazon Connect only for 24 hours.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s GetMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMetricDataInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.Filters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filters"))
	}

	if s.HistoricalMetrics == nil {
		invalidParams.Add(aws.NewErrParamRequired("HistoricalMetrics"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			invalidParams.AddNested("Filters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	// A list of HistoricalMetricResult objects, organized by Dimensions, which
	// is the ID of the resource specified in the Filters used for the request.
	// The metrics are combined with the metrics included in Collections, which
	// is a list of HisotricalMetricData objects.
	//
	// If no Grouping is specified in the request, Collections includes summary
	// data for the HistoricalMetrics.
	MetricResults []HistoricalMetricResult `type:"list"`

	// A string returned in the response. Use the value returned in the response
	// as the value of the NextToken in a subsequent request to retrieve the next
	// set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the NextToken must use the same request parameters as the
	// request that generated the token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}
