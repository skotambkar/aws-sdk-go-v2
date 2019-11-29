// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/enums"
)

type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// The duration of the events to be listed.
	Duration *int64 `type:"integer"`

	// The end time for the events to be listed.
	EndTime *time.Time `type:"timestamp"`

	// A list of event categories for the source type that you've chosen.
	EventCategories []string `type:"list"`

	// Filters applied to the action.
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The identifier of an event source.
	SourceIdentifier *string `type:"string"`

	// The type of AWS DMS resource that generates events.
	//
	// Valid values: replication-instance | replication-task
	SourceType enums.SourceType `type:"string" enum:"true"`

	// The start time for the events to be listed.
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

type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// The events described.
	Events []Event `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}