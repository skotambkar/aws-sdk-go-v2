// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/enums"
)

type DescribePublishingDestinationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the publishing destination to retrieve.
	//
	// DestinationId is a required field
	DestinationId *string `location:"uri" locationName:"destinationId" type:"string" required:"true"`

	// The unique ID of the detector associated with the publishing destination
	// to retrieve.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePublishingDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePublishingDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePublishingDestinationInput"}

	if s.DestinationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationId"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePublishingDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the publishing destination.
	//
	// DestinationId is a required field
	DestinationId *string `locationName:"destinationId" type:"string" required:"true"`

	// A DestinationProperties object that includes the DestinationArn and KmsKeyArn
	// of the publishing destination.
	//
	// DestinationProperties is a required field
	DestinationProperties *DestinationProperties `locationName:"destinationProperties" type:"structure" required:"true"`

	// The type of the publishing destination. Currently, only S3 is supported.
	//
	// DestinationType is a required field
	DestinationType enums.DestinationType `locationName:"destinationType" min:"1" type:"string" required:"true" enum:"true"`

	// The time, in epoch millisecond format, at which GuardDuty was first unable
	// to publish findings to the destination.
	//
	// PublishingFailureStartTimestamp is a required field
	PublishingFailureStartTimestamp *int64 `locationName:"publishingFailureStartTimestamp" type:"long" required:"true"`

	// The status of the publishing destination.
	//
	// Status is a required field
	Status enums.PublishingStatus `locationName:"status" min:"1" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribePublishingDestinationOutput) String() string {
	return awsutil.Prettify(s)
}
