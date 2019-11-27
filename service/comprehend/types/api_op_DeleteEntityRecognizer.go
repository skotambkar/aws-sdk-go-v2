// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteEntityRecognizerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the entity recognizer.
	//
	// EntityRecognizerArn is a required field
	EntityRecognizerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEntityRecognizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEntityRecognizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEntityRecognizerInput"}

	if s.EntityRecognizerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityRecognizerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEntityRecognizerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEntityRecognizerOutput) String() string {
	return awsutil.Prettify(s)
}
