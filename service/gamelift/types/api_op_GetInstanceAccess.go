// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type GetInstanceAccessInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet that contains the instance you want access
	// to. The fleet can be in any of the following statuses: ACTIVATING, ACTIVE,
	// or ERROR. Fleets with an ERROR status may be accessible for a short time
	// before they are deleted.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// Unique identifier for an instance you want to get access to. You can access
	// an instance in any status.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceAccessInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type GetInstanceAccessOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains connection information for a fleet instance, including
	// IP address and access credentials.
	InstanceAccess *InstanceAccess `type:"structure"`
}

// String returns the string representation
func (s GetInstanceAccessOutput) String() string {
	return awsutil.Prettify(s)
}