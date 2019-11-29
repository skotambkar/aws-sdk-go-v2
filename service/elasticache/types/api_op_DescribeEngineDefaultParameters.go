// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeEngineDefaultParameters operation.
type DescribeEngineDefaultParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache parameter group family.
	//
	// Valid values are: memcached1.4 | memcached1.5 | redis2.6 | redis2.8 | redis3.2
	// | redis4.0 | redis5.0 |
	//
	// CacheParameterGroupFamily is a required field
	CacheParameterGroupFamily *string `type:"string" required:"true"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEngineDefaultParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEngineDefaultParametersInput"}

	if s.CacheParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupFamily"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEngineDefaultParametersOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of a DescribeEngineDefaultParameters operation.
	EngineDefaults *EngineDefaults `type:"structure"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersOutput) String() string {
	return awsutil.Prettify(s)
}