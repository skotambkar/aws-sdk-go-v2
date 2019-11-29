// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The information for the launch template.
	//
	// LaunchTemplateData is a required field
	LaunchTemplateData *RequestLaunchTemplateData `type:"structure" required:"true"`

	// A name for the launch template.
	//
	// LaunchTemplateName is a required field
	LaunchTemplateName *string `min:"3" type:"string" required:"true"`

	// The tags to apply to the launch template during creation.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`

	// A description for the first version of the launch template.
	VersionDescription *string `type:"string"`
}

// String returns the string representation
func (s CreateLaunchTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLaunchTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLaunchTemplateInput"}

	if s.LaunchTemplateData == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchTemplateData"))
	}

	if s.LaunchTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchTemplateName"))
	}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}
	if s.LaunchTemplateData != nil {
		if err := s.LaunchTemplateData.Validate(); err != nil {
			invalidParams.AddNested("LaunchTemplateData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLaunchTemplateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template.
	LaunchTemplate *LaunchTemplate `locationName:"launchTemplate" type:"structure"`
}

// String returns the string representation
func (s CreateLaunchTemplateOutput) String() string {
	return awsutil.Prettify(s)
}