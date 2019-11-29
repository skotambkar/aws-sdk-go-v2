// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type InvokeDeviceMethodInput struct {
	_ struct{} `type:"structure"`

	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	// The device method to invoke.
	DeviceMethod *DeviceMethod `locationName:"deviceMethod" type:"structure"`

	// A JSON encoded string containing the device method request parameters.
	DeviceMethodParameters *string `locationName:"deviceMethodParameters" type:"string"`
}

// String returns the string representation
func (s InvokeDeviceMethodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeDeviceMethodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InvokeDeviceMethodInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InvokeDeviceMethodOutput struct {
	_ struct{} `type:"structure"`

	// A JSON encoded string containing the device method response.
	DeviceMethodResponse *string `locationName:"deviceMethodResponse" type:"string"`
}

// String returns the string representation
func (s InvokeDeviceMethodOutput) String() string {
	return awsutil.Prettify(s)
}