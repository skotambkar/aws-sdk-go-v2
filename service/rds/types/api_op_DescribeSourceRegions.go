// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSourceRegionsInput struct {
	_ struct{} `type:"structure"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeSourceRegions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The source AWS Region name. For example, us-east-1.
	//
	// Constraints:
	//
	//    * Must specify a valid AWS Region name.
	RegionName *string `type:"string"`
}

// String returns the string representation
func (s DescribeSourceRegionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSourceRegionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSourceRegionsInput"}
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

// Contains the result of a successful invocation of the DescribeSourceRegions
// action.
type DescribeSourceRegionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of SourceRegion instances that contains each source AWS Region that
	// the current AWS Region can get a Read Replica or a DB snapshot from.
	SourceRegions []SourceRegion `locationNameList:"SourceRegion" type:"list"`
}

// String returns the string representation
func (s DescribeSourceRegionsOutput) String() string {
	return awsutil.Prettify(s)
}
