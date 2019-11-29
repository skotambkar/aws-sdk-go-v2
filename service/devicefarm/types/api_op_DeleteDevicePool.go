// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the delete device pool operation.
type DeleteDevicePoolInput struct {
	_ struct{} `type:"structure"`

	// Represents the Amazon Resource Name (ARN) of the Device Farm device pool
	// you wish to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDevicePoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDevicePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDevicePoolInput"}

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

// Represents the result of a delete device pool request.
type DeleteDevicePoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDevicePoolOutput) String() string {
	return awsutil.Prettify(s)
}