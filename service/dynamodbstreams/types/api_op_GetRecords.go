// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetRecords operation.
type GetRecordsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of records to return from the shard. The upper limit is
	// 1000.
	Limit *int64 `min:"1" type:"integer"`

	// A shard iterator that was retrieved from a previous GetShardIterator operation.
	// This iterator can be used to access the stream records in this shard.
	//
	// ShardIterator is a required field
	ShardIterator *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRecordsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.ShardIterator == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardIterator"))
	}
	if s.ShardIterator != nil && len(*s.ShardIterator) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardIterator", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetRecords operation.
type GetRecordsOutput struct {
	_ struct{} `type:"structure"`

	// The next position in the shard from which to start sequentially reading stream
	// records. If set to null, the shard has been closed and the requested iterator
	// will not return any more data.
	NextShardIterator *string `min:"1" type:"string"`

	// The stream records from the shard, which were retrieved using the shard iterator.
	Records []Record `type:"list"`
}

// String returns the string representation
func (s GetRecordsOutput) String() string {
	return awsutil.Prettify(s)
}