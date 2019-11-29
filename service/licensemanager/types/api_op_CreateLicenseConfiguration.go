// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/enums"
)

type CreateLicenseConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Human-friendly description of the license configuration.
	Description *string `type:"string"`

	// Number of licenses managed by the license configuration.
	LicenseCount *int64 `type:"long"`

	// Flag indicating whether hard or soft license enforcement is used. Exceeding
	// a hard limit results in the blocked deployment of new instances.
	LicenseCountHardLimit *bool `type:"boolean"`

	// Dimension to use to track the license inventory.
	//
	// LicenseCountingType is a required field
	LicenseCountingType enums.LicenseCountingType `type:"string" required:"true" enum:"true"`

	// Array of configured License Manager rules.
	LicenseRules []string `type:"list"`

	// Name of the license configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The tags to apply to the resources during launch. You can only tag instances
	// and volumes on launch. The specified tags are applied to all instances or
	// volumes that are created during launch. To tag a resource after it has been
	// created, see CreateTags .
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateLicenseConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLicenseConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLicenseConfigurationInput"}
	if len(s.LicenseCountingType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LicenseCountingType"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLicenseConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the license configuration object after its creation.
	LicenseConfigurationArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLicenseConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}