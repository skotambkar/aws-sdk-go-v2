// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReservedDBInstancesInput struct {
	_ struct{} `type:"structure"`

	// The DB instance class filter value. Specify this parameter to show only those
	// reservations matching the specified DB instances class.
	DBInstanceClass *string `type:"string"`

	// The duration filter value, specified in years or seconds. Specify this parameter
	// to show only reservations for this duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string `type:"string"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The lease identifier filter value. Specify this parameter to show only the
	// reservation that matches the specified lease ID.
	//
	// AWS Support might request the lease ID for an issue related to a reserved
	// DB instance.
	LeaseId *string `type:"string"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included
	// in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A value that indicates whether to show only those reservations that support
	// Multi-AZ.
	MultiAZ *bool `type:"boolean"`

	// The offering type filter value. Specify this parameter to show only the available
	// offerings matching the specified offering type.
	//
	// Valid Values: "Partial Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string `type:"string"`

	// The product description filter value. Specify this parameter to show only
	// those reservations matching the specified product description.
	ProductDescription *string `type:"string"`

	// The reserved DB instance identifier filter value. Specify this parameter
	// to show only the reservation that matches the specified reservation ID.
	ReservedDBInstanceId *string `type:"string"`

	// The offering identifier filter value. Specify this parameter to show only
	// purchased reservations matching the specified offering identifier.
	ReservedDBInstancesOfferingId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedDBInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReservedDBInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReservedDBInstancesInput"}
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

// Contains the result of a successful invocation of the DescribeReservedDBInstances
// action.
type DescribeReservedDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of reserved DB instances.
	ReservedDBInstances []ReservedDBInstance `locationNameList:"ReservedDBInstance" type:"list"`
}

// String returns the string representation
func (s DescribeReservedDBInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
