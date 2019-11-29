// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForDeliveryStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream whose tags you want to list.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// The key to use as the starting point for the list of tags. If you set this
	// parameter, ListTagsForDeliveryStream gets all tags that occur after ExclusiveStartTagKey.
	ExclusiveStartTagKey *string `min:"1" type:"string"`

	// The number of tags to return. If this number is less than the total number
	// of tags associated with the delivery stream, HasMoreTags is set to true in
	// the response. To list additional tags, set ExclusiveStartTagKey to the last
	// key in the response.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s ListTagsForDeliveryStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForDeliveryStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForDeliveryStreamInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}
	if s.ExclusiveStartTagKey != nil && len(*s.ExclusiveStartTagKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartTagKey", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForDeliveryStreamOutput struct {
	_ struct{} `type:"structure"`

	// If this is true in the response, more tags are available. To list the remaining
	// tags, set ExclusiveStartTagKey to the key of the last tag returned and call
	// ListTagsForDeliveryStream again.
	//
	// HasMoreTags is a required field
	HasMoreTags *bool `type:"boolean" required:"true"`

	// A list of tags associated with DeliveryStreamName, starting with the first
	// tag after ExclusiveStartTagKey and up to the specified Limit.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForDeliveryStreamOutput) String() string {
	return awsutil.Prettify(s)
}