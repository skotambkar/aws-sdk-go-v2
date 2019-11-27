// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/rds/enums"
)

type StopActivityStreamInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether or not the database activity stream is to stop as soon
	// as possible, regardless of the maintenance window for the database.
	ApplyImmediately *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the DB cluster for the database activity
	// stream. For example, arn:aws:rds:us-east-1:12345667890:cluster:das-cluster.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopActivityStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopActivityStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopActivityStreamInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopActivityStreamOutput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon Kinesis data stream used for the database activity
	// stream.
	KinesisStreamName *string `type:"string"`

	// The AWS KMS key identifier used for encrypting messages in the database activity
	// stream.
	KmsKeyId *string `type:"string"`

	// The status of the database activity stream.
	Status enums.ActivityStreamStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StopActivityStreamOutput) String() string {
	return awsutil.Prettify(s)
}
