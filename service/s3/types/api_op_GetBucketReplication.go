// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBucketReplicationInput struct {
	_ struct{} `type:"structure"`

	// The bucket name for which to get the replication information.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketReplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketReplicationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketReplicationInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketReplicationOutput struct {
	_ struct{} `type:"structure" payload:"ReplicationConfiguration"`

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	ReplicationConfiguration *ReplicationConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketReplicationOutput) String() string {
	return awsutil.Prettify(s)
}