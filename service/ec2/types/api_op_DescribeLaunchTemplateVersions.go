// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLaunchTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * create-time - The time the launch template version was created.
	//
	//    * ebs-optimized - A boolean that indicates whether the instance is optimized
	//    for Amazon EBS I/O.
	//
	//    * iam-instance-profile - The ARN of the IAM instance profile.
	//
	//    * image-id - The ID of the AMI.
	//
	//    * instance-type - The instance type.
	//
	//    * is-default-version - A boolean that indicates whether the launch template
	//    version is the default version.
	//
	//    * kernel-id - The kernel ID.
	//
	//    * ram-disk-id - The RAM disk ID.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	MaxResults *int64 `type:"integer"`

	// The version number up to which to describe launch template versions.
	MaxVersion *string `type:"string"`

	// The version number after which to describe launch template versions.
	MinVersion *string `type:"string"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`

	// One or more versions of the launch template.
	Versions []string `locationName:"LaunchTemplateVersion" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLaunchTemplateVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLaunchTemplateVersionsInput"}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template versions.
	LaunchTemplateVersions []LaunchTemplateVersion `locationName:"launchTemplateVersionSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsOutput) String() string {
	return awsutil.Prettify(s)
}
