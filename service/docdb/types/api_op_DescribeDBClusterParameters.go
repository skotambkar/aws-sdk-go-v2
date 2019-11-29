// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to DescribeDBClusterParameters.
type DescribeDBClusterParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB cluster parameter group to return parameter details
	// for.
	//
	// Constraints:
	//
	//    * If provided, must match the name of an existing DBClusterParameterGroup.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

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

	// A value that indicates to return only parameters for a specific source. Parameter
	// sources can be engine, service, or customer.
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterParametersInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
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

// Represents the output of DBClusterParameterGroup.
type DescribeDBClusterParametersOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// Provides a list of parameters for the DB cluster parameter group.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s DescribeDBClusterParametersOutput) String() string {
	return awsutil.Prettify(s)
}
