// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOptionGroupsInput struct {
	_ struct{} `type:"structure"`

	// Filters the list of option groups to only include groups associated with
	// a specific database engine.
	EngineName *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// Filters the list of option groups to only include groups associated with
	// a specific database engine version. If specified, then EngineName must also
	// be specified.
	MajorEngineVersion *string `type:"string"`

	// An optional pagination token provided by a previous DescribeOptionGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the option group to describe. Can't be supplied together with
	// EngineName or MajorEngineVersion.
	OptionGroupName *string `type:"string"`
}

// String returns the string representation
func (s DescribeOptionGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOptionGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOptionGroupsInput"}
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

// List of option groups.
type DescribeOptionGroupsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// List of option groups.
	OptionGroupsList []OptionGroup `locationNameList:"OptionGroup" type:"list"`
}

// String returns the string representation
func (s DescribeOptionGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
