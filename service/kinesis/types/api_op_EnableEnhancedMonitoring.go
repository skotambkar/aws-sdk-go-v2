// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/enums"
)

// Represents the input for EnableEnhancedMonitoring.
type EnableEnhancedMonitoringInput struct {
	_ struct{} `type:"structure"`

	// List of shard-level metrics to enable.
	//
	// The following are the valid shard-level metrics. The value "ALL" enables
	// every metric.
	//
	//    * IncomingBytes
	//
	//    * IncomingRecords
	//
	//    * OutgoingBytes
	//
	//    * OutgoingRecords
	//
	//    * WriteProvisionedThroughputExceeded
	//
	//    * ReadProvisionedThroughputExceeded
	//
	//    * IteratorAgeMilliseconds
	//
	//    * ALL
	//
	// For more information, see Monitoring the Amazon Kinesis Data Streams Service
	// with Amazon CloudWatch (http://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	//
	// ShardLevelMetrics is a required field
	ShardLevelMetrics []enums.MetricsName `min:"1" type:"list" required:"true"`

	// The name of the stream for which to enable enhanced monitoring.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableEnhancedMonitoringInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableEnhancedMonitoringInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableEnhancedMonitoringInput"}

	if s.ShardLevelMetrics == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardLevelMetrics"))
	}
	if s.ShardLevelMetrics != nil && len(s.ShardLevelMetrics) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardLevelMetrics", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output for EnableEnhancedMonitoring and DisableEnhancedMonitoring.
type EnableEnhancedMonitoringOutput struct {
	_ struct{} `type:"structure"`

	// Represents the current state of the metrics that are in the enhanced state
	// before the operation.
	CurrentShardLevelMetrics []enums.MetricsName `min:"1" type:"list"`

	// Represents the list of all the metrics that would be in the enhanced state
	// after the operation.
	DesiredShardLevelMetrics []enums.MetricsName `min:"1" type:"list"`

	// The name of the Kinesis data stream.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EnableEnhancedMonitoringOutput) String() string {
	return awsutil.Prettify(s)
}