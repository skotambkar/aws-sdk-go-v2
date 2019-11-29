// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDeviceEventsInput struct {
	_ struct{} `type:"structure"`

	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	// FromTimeStamp is a required field
	FromTimeStamp *time.Time `location:"querystring" locationName:"fromTimeStamp" type:"timestamp" timestampFormat:"iso8601" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// ToTimeStamp is a required field
	ToTimeStamp *time.Time `location:"querystring" locationName:"toTimeStamp" type:"timestamp" timestampFormat:"iso8601" required:"true"`
}

// String returns the string representation
func (s ListDeviceEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeviceEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeviceEventsInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if s.FromTimeStamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromTimeStamp"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ToTimeStamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ToTimeStamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDeviceEventsOutput struct {
	_ struct{} `type:"structure"`

	Events []DeviceEvent `locationName:"events" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeviceEventsOutput) String() string {
	return awsutil.Prettify(s)
}