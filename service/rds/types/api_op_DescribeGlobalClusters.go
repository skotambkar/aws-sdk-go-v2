// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeGlobalClustersInput struct {
	_ struct{} `type:"structure"`

	// A filter that specifies one or more global DB clusters to describe.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include information
	//    about the DB clusters identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match an existing DBClusterIdentifier.
	GlobalClusterIdentifier *string `type:"string"`

	// An optional pagination token provided by a previous DescribeGlobalClusters
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
}

// String returns the string representation
func (s DescribeGlobalClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGlobalClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGlobalClustersInput"}
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

type DescribeGlobalClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of global clusters returned by this request.
	GlobalClusters []GlobalCluster `locationNameList:"GlobalClusterMember" type:"list"`

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeGlobalClustersOutput) String() string {
	return awsutil.Prettify(s)
}
