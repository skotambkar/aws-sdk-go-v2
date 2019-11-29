// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UntagInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource from which to remove tags.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"Arn" min:"12" type:"string" required:"true"`

	// The keys of the tags to be removed.
	//
	// Keys is a required field
	Keys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s UntagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 12))
	}

	if s.Keys == nil {
		invalidParams.Add(aws.NewErrParamRequired("Keys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource from which tags have been removed.
	Arn *string `min:"12" type:"string"`

	// The keys of tags that have been removed.
	Keys []string `type:"list"`
}

// String returns the string representation
func (s UntagOutput) String() string {
	return awsutil.Prettify(s)
}
