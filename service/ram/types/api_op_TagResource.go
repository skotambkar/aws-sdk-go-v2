// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string" required:"true"`

	// One or more tags.
	//
	// Tags is a required field
	Tags []Tag `locationName:"tags" type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}