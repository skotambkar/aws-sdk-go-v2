// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the delete upload operation.
type DeleteUploadInput struct {
	_ struct{} `type:"structure"`

	// Represents the Amazon Resource Name (ARN) of the Device Farm upload you wish
	// to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUploadInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a delete upload request.
type DeleteUploadOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUploadOutput) String() string {
	return awsutil.Prettify(s)
}