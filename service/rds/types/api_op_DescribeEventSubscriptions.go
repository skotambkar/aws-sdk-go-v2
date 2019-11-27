// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the RDS event notification subscription you want to describe.
	SubscriptionName *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventSubscriptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventSubscriptionsInput"}
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

// Data returned by the DescribeEventSubscriptions action.
type DescribeEventSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of EventSubscriptions data types.
	EventSubscriptionsList []EventSubscription `locationNameList:"EventSubscription" type:"list"`

	// An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}
