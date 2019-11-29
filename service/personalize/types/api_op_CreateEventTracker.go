// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEventTrackerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset group that receives the event
	// data.
	//
	// DatasetGroupArn is a required field
	DatasetGroupArn *string `locationName:"datasetGroupArn" type:"string" required:"true"`

	// The name for the event tracker.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEventTrackerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEventTrackerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEventTrackerInput"}

	if s.DatasetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetGroupArn"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateEventTrackerOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the event tracker.
	EventTrackerArn *string `locationName:"eventTrackerArn" type:"string"`

	// The ID of the event tracker. Include this ID in requests to the PutEvents
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html)
	// API.
	TrackingId *string `locationName:"trackingId" type:"string"`
}

// String returns the string representation
func (s CreateEventTrackerOutput) String() string {
	return awsutil.Prettify(s)
}