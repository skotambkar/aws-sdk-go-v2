// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDatastoreInput struct {
	_ struct{} `type:"structure"`

	// The name of the data store
	//
	// DatastoreName is a required field
	DatastoreName *string `location:"uri" locationName:"datastoreName" min:"1" type:"string" required:"true"`

	// If true, additional statistical information about the data store is included
	// in the response. This feature cannot be used with a data store whose S3 storage
	// is customer-managed.
	IncludeStatistics *bool `location:"querystring" locationName:"includeStatistics" type:"boolean"`
}

// String returns the string representation
func (s DescribeDatastoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatastoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatastoreInput"}

	if s.DatastoreName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatastoreName"))
	}
	if s.DatastoreName != nil && len(*s.DatastoreName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatastoreName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDatastoreOutput struct {
	_ struct{} `type:"structure"`

	// Information about the data store.
	Datastore *Datastore `locationName:"datastore" type:"structure"`

	// Additional statistical information about the data store. Included if the
	// 'includeStatistics' parameter is set to true in the request.
	Statistics *DatastoreStatistics `locationName:"statistics" type:"structure"`
}

// String returns the string representation
func (s DescribeDatastoreOutput) String() string {
	return awsutil.Prettify(s)
}