// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLaunchTemplateVersionInput struct {
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

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`

	// The version number of the launch template version on which to base the new
	// version. The new version inherits the same launch parameters as the source
	// version, except for parameters that you specify in LaunchTemplateData. Snapshots
	// applied to the block device mapping are ignored when creating a new version
	// unless they are explicitly included.
	SourceVersion *string `type:"string"`

	// A description for the version of the launch template.
	VersionDescription *string `type:"string"`
}

// String returns the string representation
func (s CreateLaunchTemplateVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLaunchTemplateVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLaunchTemplateVersionInput"}

	if s.LaunchTemplateData == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchTemplateData"))
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

type CreateLaunchTemplateVersionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template version.
	LaunchTemplateVersion *LaunchTemplateVersion `locationName:"launchTemplateVersion" type:"structure"`
}

// String returns the string representation
func (s CreateLaunchTemplateVersionOutput) String() string {
	return awsutil.Prettify(s)
}