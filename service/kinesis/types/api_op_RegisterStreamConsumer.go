// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterStreamConsumerInput struct {
	_ struct{} `type:"structure"`

	// For a given Kinesis data stream, each consumer must have a unique name. However,
	// consumer names don't have to be unique across data streams.
	//
	// ConsumerName is a required field
	ConsumerName *string `min:"1" type:"string" required:"true"`

	// The ARN of the Kinesis data stream that you want to register the consumer
	// with. For more info, see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterStreamConsumerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterStreamConsumerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterStreamConsumerInput"}

	if s.ConsumerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConsumerName"))
	}
	if s.ConsumerName != nil && len(*s.ConsumerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConsumerName", 1))
	}

	if s.StreamARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamARN"))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterStreamConsumerOutput struct {
	_ struct{} `type:"structure"`

	// An object that represents the details of the consumer you registered. When
	// you register a consumer, it gets an ARN that is generated by Kinesis Data
	// Streams.
	//
	// Consumer is a required field
	Consumer *Consumer `type:"structure" required:"true"`
}

// String returns the string representation
func (s RegisterStreamConsumerOutput) String() string {
	return awsutil.Prettify(s)
}