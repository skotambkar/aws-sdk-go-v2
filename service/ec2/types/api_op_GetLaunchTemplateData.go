// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLaunchTemplateDataInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetLaunchTemplateDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLaunchTemplateDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLaunchTemplateDataInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLaunchTemplateDataOutput struct {
	_ struct{} `type:"structure"`

	// The instance data.
	LaunchTemplateData *ResponseLaunchTemplateData `locationName:"launchTemplateData" type:"structure"`
}

// String returns the string representation
func (s GetLaunchTemplateDataOutput) String() string {
	return awsutil.Prettify(s)
}
