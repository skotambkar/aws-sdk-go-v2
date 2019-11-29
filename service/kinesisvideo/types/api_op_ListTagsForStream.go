// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForStreamInput struct {
	_ struct{} `type:"structure"`

	// If you specify this parameter and the result of a ListTagsForStream call
	// is truncated, the response includes a token that you can use in the next
	// request to fetch the next batch of tags.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) of the stream that you want to list tags for.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream that you want to list tags for.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListTagsForStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForStreamInput"}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForStreamOutput struct {
	_ struct{} `type:"structure"`

	// If you specify this parameter and the result of a ListTags call is truncated,
	// the response includes a token that you can use in the next request to fetch
	// the next set of tags.
	NextToken *string `type:"string"`

	// A map of tag keys and values associated with the specified stream.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s ListTagsForStreamOutput) String() string {
	return awsutil.Prettify(s)
}