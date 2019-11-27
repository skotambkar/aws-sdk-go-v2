// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMLTransformInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the transform to delete.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMLTransformInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMLTransformInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMLTransformInput"}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMLTransformOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the transform that was deleted.
	TransformId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteMLTransformOutput) String() string {
	return awsutil.Prettify(s)
}
