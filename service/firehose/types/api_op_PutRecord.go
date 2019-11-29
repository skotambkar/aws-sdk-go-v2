// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutRecordInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// The record.
	//
	// Record is a required field
	Record *Record `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutRecordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecordInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if s.Record == nil {
		invalidParams.Add(aws.NewErrParamRequired("Record"))
	}
	if s.Record != nil {
		if err := s.Record.Validate(); err != nil {
			invalidParams.AddNested("Record", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutRecordOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether server-side encryption (SSE) was enabled during this operation.
	Encrypted *bool `type:"boolean"`

	// The ID of the record.
	//
	// RecordId is a required field
	RecordId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecordOutput) String() string {
	return awsutil.Prettify(s)
}