// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListStreamConsumersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of consumers that you want a single call of ListStreamConsumers
	// to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// When the number of consumers that are registered with the data stream is
	// greater than the default value for the MaxResults parameter, or if you explicitly
	// specify a value for MaxResults that is less than the number of consumers
	// that are registered with the data stream, the response includes a pagination
	// token named NextToken. You can specify this NextToken value in a subsequent
	// call to ListStreamConsumers to list the next set of registered consumers.
	//
	// Don't specify StreamName or StreamCreationTimestamp if you specify NextToken
	// because the latter unambiguously identifies the stream.
	//
	// You can optionally specify a value for the MaxResults parameter when you
	// specify NextToken. If you specify a MaxResults value that is less than the
	// number of consumers that the operation returns if you don't specify MaxResults,
	// the response will contain a new NextToken value. You can use the new NextToken
	// value in a subsequent call to the ListStreamConsumers operation to list the
	// next set of consumers.
	//
	// Tokens expire after 300 seconds. When you obtain a value for NextToken in
	// the response to a call to ListStreamConsumers, you have 300 seconds to use
	// that value. If you specify an expired token in a call to ListStreamConsumers,
	// you get ExpiredNextTokenException.
	NextToken *string `min:"1" type:"string"`

	// The ARN of the Kinesis data stream for which you want to list the registered
	// consumers. For more information, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`

	// Specify this input parameter to distinguish data streams that have the same
	// name. For example, if you create a data stream and then delete it, and you
	// later create another data stream with the same name, you can use this input
	// parameter to specify which of the two streams you want to list the consumers
	// for.
	//
	// You can't specify this parameter if you specify the NextToken parameter.
	StreamCreationTimestamp *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ListStreamConsumersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamConsumersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamConsumersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
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

type ListStreamConsumersOutput struct {
	_ struct{} `type:"structure"`

	// An array of JSON objects. Each object represents one registered consumer.
	Consumers []Consumer `type:"list"`

	// When the number of consumers that are registered with the data stream is
	// greater than the default value for the MaxResults parameter, or if you explicitly
	// specify a value for MaxResults that is less than the number of registered
	// consumers, the response includes a pagination token named NextToken. You
	// can specify this NextToken value in a subsequent call to ListStreamConsumers
	// to list the next set of registered consumers. For more information about
	// the use of this pagination token when calling the ListStreamConsumers operation,
	// see ListStreamConsumersInput$NextToken.
	//
	// Tokens expire after 300 seconds. When you obtain a value for NextToken in
	// the response to a call to ListStreamConsumers, you have 300 seconds to use
	// that value. If you specify an expired token in a call to ListStreamConsumers,
	// you get ExpiredNextTokenException.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListStreamConsumersOutput) String() string {
	return awsutil.Prettify(s)
}