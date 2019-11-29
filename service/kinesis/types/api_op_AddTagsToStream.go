// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for AddTagsToStream.
type AddTagsToStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the stream.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// A set of up to 10 key-value pairs to use to create the tags.
	//
	// Tags is a required field
	Tags map[string]string `min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s AddTagsToStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToStreamInput"}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddTagsToStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToStreamOutput) String() string {
	return awsutil.Prettify(s)
}
