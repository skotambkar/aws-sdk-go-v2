// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartDeliveryStreamEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream for which you want to enable server-side
	// encryption (SSE).
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartDeliveryStreamEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDeliveryStreamEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDeliveryStreamEncryptionInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartDeliveryStreamEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartDeliveryStreamEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}
