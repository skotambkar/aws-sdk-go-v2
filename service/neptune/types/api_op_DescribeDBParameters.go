// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB parameter group to return details for.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The parameter types to return.
	//
	// Default: All parameter types returned
	//
	// Valid Values: user | system | engine-default
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBParametersInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}
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

type DescribeDBParametersOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of Parameter values.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s DescribeDBParametersOutput) String() string {
	return awsutil.Prettify(s)
}
