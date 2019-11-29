// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the create device pool operation.
type CreateDevicePoolInput struct {
	_ struct{} `type:"structure"`

	// The device pool's description.
	Description *string `locationName:"description" type:"string"`

	// The number of devices that Device Farm can add to your device pool. Device
	// Farm adds devices that are available and that meet the criteria that you
	// assign for the rules parameter. Depending on how many devices meet these
	// constraints, your device pool might contain fewer devices than the value
	// for this parameter.
	//
	// By specifying the maximum number of devices, you can control the costs that
	// you incur by running tests.
	MaxDevices *int64 `locationName:"maxDevices" type:"integer"`

	// The device pool's name.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The ARN of the project for the device pool.
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"32" type:"string" required:"true"`

	// The device pool's rules.
	//
	// Rules is a required field
	Rules []Rule `locationName:"rules" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateDevicePoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDevicePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDevicePoolInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 32))
	}

	if s.Rules == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rules"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a create device pool request.
type CreateDevicePoolOutput struct {
	_ struct{} `type:"structure"`

	// The newly created device pool.
	DevicePool *DevicePool `locationName:"devicePool" type:"structure"`
}

// String returns the string representation
func (s CreateDevicePoolOutput) String() string {
	return awsutil.Prettify(s)
}