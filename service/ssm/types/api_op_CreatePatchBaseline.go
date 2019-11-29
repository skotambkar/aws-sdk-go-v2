// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type CreatePatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// A set of rules used to include patches in the baseline.
	ApprovalRules *PatchRuleGroup `type:"structure"`

	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see Package Name Formats for Approved and Rejected Patch
	// Lists (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	ApprovedPatches []string `type:"list"`

	// Defines the compliance level for approved patches. This means that if an
	// approved patch is reported as missing, this is the severity of the compliance
	// violation. The default value is UNSPECIFIED.
	ApprovedPatchesComplianceLevel enums.PatchComplianceLevel `type:"string" enum:"true"`

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is 'false'. Applies
	// to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `type:"boolean"`

	// User-provided idempotency token.
	ClientToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// A description of the patch baseline.
	Description *string `min:"1" type:"string"`

	// A set of global filters used to include patches in the baseline.
	GlobalFilters *PatchFilterGroup `type:"structure"`

	// The name of the patch baseline.
	//
	// Name is a required field
	Name *string `min:"3" type:"string" required:"true"`

	// Defines the operating system the patch baseline applies to. The Default value
	// is WINDOWS.
	OperatingSystem enums.OperatingSystem `type:"string" enum:"true"`

	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see Package Name Formats for Approved and Rejected Patch
	// Lists (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	RejectedPatches []string `type:"list"`

	// The action for Patch Manager to take on patches included in the RejectedPackages
	// list.
	//
	//    * ALLOW_AS_DEPENDENCY: A package in the Rejected patches list is installed
	//    only if it is a dependency of another package. It is considered compliant
	//    with the patch baseline, and its status is reported as InstalledOther.
	//    This is the default action if no option is specified.
	//
	//    * BLOCK: Packages in the RejectedPatches list, and packages that include
	//    them as dependencies, are not installed under any circumstances. If a
	//    package was installed before it was added to the Rejected patches list,
	//    it is considered non-compliant with the patch baseline, and its status
	//    is reported as InstalledRejected.
	RejectedPatchesAction enums.PatchAction `type:"string" enum:"true"`

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []PatchSource `type:"list"`

	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	// For example, you might want to tag a patch baseline to identify the severity
	// level of patches it specifies and the operating system family it applies
	// to. In this case, you could specify the following key name/value pairs:
	//
	//    * Key=PatchSeverity,Value=Critical
	//
	//    * Key=OS,Value=Windows
	//
	// To add tags to an existing patch baseline, use the AddTagsToResource action.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePatchBaselineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePatchBaselineInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}
	if s.ApprovalRules != nil {
		if err := s.ApprovalRules.Validate(); err != nil {
			invalidParams.AddNested("ApprovalRules", err.(aws.ErrInvalidParams))
		}
	}
	if s.GlobalFilters != nil {
		if err := s.GlobalFilters.Validate(); err != nil {
			invalidParams.AddNested("GlobalFilters", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sources != nil {
		for i, v := range s.Sources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Sources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the created patch baseline.
	BaselineId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreatePatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}