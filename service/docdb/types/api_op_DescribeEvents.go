// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/docdb/enums"
)

// Represents the input to DescribeEvents.
type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// The number of minutes to retrieve events for.
	//
	// Default: 60
	Duration *int64 `type:"integer"`

	// The end of the time interval for which to retrieve events, specified in ISO
	// 8601 format.
	//
	// Example: 2009-07-08T18:00Z
	EndTime *time.Time `type:"timestamp"`

	// A list of event categories that trigger notifications for an event notification
	// subscription.
	EventCategories []string `locationNameList:"EventCategory" type:"list"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token (marker) is
	// included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The identifier of the event source for which events are returned. If not
	// specified, then all sources are included in the response.
	//
	// Constraints:
	//
	//    * If SourceIdentifier is provided, SourceType must also be provided.
	//
	//    * If the source type is DBInstance, a DBInstanceIdentifier must be provided.
	//
	//    * If the source type is DBSecurityGroup, a DBSecurityGroupName must be
	//    provided.
	//
	//    * If the source type is DBParameterGroup, a DBParameterGroupName must
	//    be provided.
	//
	//    * If the source type is DBSnapshot, a DBSnapshotIdentifier must be provided.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	SourceIdentifier *string `type:"string"`

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType enums.SourceType `type:"string" enum:"true"`

	// The beginning of the time interval to retrieve events for, specified in ISO
	// 8601 format.
	//
	// Example: 2009-07-08T18:00Z
	StartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of DescribeEvents.
type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about one or more events.
	Events []Event `locationNameList:"Event" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}