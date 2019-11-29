// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDeliveryStreamInput struct {
	_ struct{} `type:"structure"`

	// Set this to true if you want to delete the delivery stream even if Kinesis
	// Data Firehose is unable to retire the grant for the CMK. Kinesis Data Firehose
	// might be unable to retire the grant due to a customer error, such as when
	// the CMK or the grant are in an invalid state. If you force deletion, you
	// can then use the RevokeGrant (https://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html)
	// operation to revoke the grant you gave to Kinesis Data Firehose. If a failure
	// to retire the grant happens due to an AWS KMS issue, Kinesis Data Firehose
	// keeps retrying the delete operation.
	//
	// The default value is false.
	AllowForceDelete *bool `type:"boolean"`

	// The name of the delivery stream.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDeliveryStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDeliveryStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDeliveryStreamInput"}

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

type DeleteDeliveryStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDeliveryStreamOutput) String() string {
	return awsutil.Prettify(s)
}