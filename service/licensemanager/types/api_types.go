// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/enums"
)

// Details about license consumption.
type ConsumedLicenseSummary struct {
	_ struct{} `type:"structure"`

	// Number of licenses consumed by a resource.
	ConsumedLicenses *int64 `type:"long"`

	// Resource type of the resource consuming a license (instance, host, or AMI).
	ResourceType enums.ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ConsumedLicenseSummary) String() string {
	return awsutil.Prettify(s)
}

// A filter name and value pair that is used to return a more specific list
// of results from a describe operation. Filters can be used to match a set
// of resources by specific criteria, such as tags, attributes, or IDs. The
// filters supported by a Describe operation are documented with the Describe
// operation.
type Filter struct {
	_ struct{} `type:"structure"`

	// Name of the filter. Filter names are case-sensitive.
	Name *string `type:"string"`

	// One or more filter values. Filter values are case-sensitive.
	Values []string `type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// An inventory filter object.
type InventoryFilter struct {
	_ struct{} `type:"structure"`

	// The condition of the filter.
	//
	// Condition is a required field
	Condition enums.InventoryFilterCondition `type:"string" required:"true" enum:"true"`

	// The name of the filter.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Value of the filter.
	Value *string `type:"string"`
}

// String returns the string representation
func (s InventoryFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InventoryFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InventoryFilter"}
	if len(s.Condition) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Condition"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A license configuration is an abstraction of a customer license agreement
// that can be consumed and enforced by License Manager. Components include
// specifications for the license type (licensing by instance, socket, CPU,
// or VCPU), tenancy (shared tenancy, Amazon EC2 Dedicated Instance, Amazon
// EC2 Dedicated Host, or any of these), host affinity (how long a VM must be
// associated with a host), the number of licenses purchased and used.
type LicenseConfiguration struct {
	_ struct{} `type:"structure"`

	// List of summaries for licenses consumed by various resources.
	ConsumedLicenseSummaryList []ConsumedLicenseSummary `type:"list"`

	// Number of licenses consumed.
	ConsumedLicenses *int64 `type:"long"`

	// Description of the license configuration.
	Description *string `type:"string"`

	// ARN of the LicenseConfiguration object.
	LicenseConfigurationArn *string `type:"string"`

	// Unique ID of the LicenseConfiguration object.
	LicenseConfigurationId *string `type:"string"`

	// Number of licenses managed by the license configuration.
	LicenseCount *int64 `type:"long"`

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `type:"boolean"`

	// Dimension to use to track license inventory.
	LicenseCountingType enums.LicenseCountingType `type:"string" enum:"true"`

	// Array of configured License Manager rules.
	LicenseRules []string `type:"list"`

	// List of summaries for managed resources.
	ManagedResourceSummaryList []ManagedResourceSummary `type:"list"`

	// Name of the license configuration.
	Name *string `type:"string"`

	// Account ID of the license configuration's owner.
	OwnerAccountId *string `type:"string"`

	// Status of the license configuration.
	Status *string `type:"string"`
}

// String returns the string representation
func (s LicenseConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Describes a server resource that is associated with a license configuration.
type LicenseConfigurationAssociation struct {
	_ struct{} `type:"structure"`

	// Time when the license configuration was associated with the resource.
	AssociationTime *time.Time `type:"timestamp"`

	// ARN of the resource associated with the license configuration.
	ResourceArn *string `type:"string"`

	// ID of the AWS account that owns the resource consuming licenses.
	ResourceOwnerId *string `type:"string"`

	// Type of server resource.
	ResourceType enums.ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s LicenseConfigurationAssociation) String() string {
	return awsutil.Prettify(s)
}

// Contains details of the usage of each resource from the license pool.
type LicenseConfigurationUsage struct {
	_ struct{} `type:"structure"`

	// Time when the license configuration was initially associated with a resource.
	AssociationTime *time.Time `type:"timestamp"`

	// Number of licenses consumed out of the total provisioned in the license configuration.
	ConsumedLicenses *int64 `type:"long"`

	// ARN of the resource associated with a license configuration.
	ResourceArn *string `type:"string"`

	// ID of the account that owns a resource that is associated with the license
	// configuration.
	ResourceOwnerId *string `type:"string"`

	// Status of a resource associated with the license configuration.
	ResourceStatus *string `type:"string"`

	// Type of resource associated with athe license configuration.
	ResourceType enums.ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s LicenseConfigurationUsage) String() string {
	return awsutil.Prettify(s)
}

// Object used for associating a license configuration with a resource.
type LicenseSpecification struct {
	_ struct{} `type:"structure"`

	// ARN of the LicenseConfiguration object.
	//
	// LicenseConfigurationArn is a required field
	LicenseConfigurationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s LicenseSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LicenseSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LicenseSpecification"}

	if s.LicenseConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LicenseConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Summary for a resource.
type ManagedResourceSummary struct {
	_ struct{} `type:"structure"`

	// Number of resources associated with licenses.
	AssociationCount *int64 `type:"long"`

	// Type of resource associated with a license (instance, host, or AMI).
	ResourceType enums.ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ManagedResourceSummary) String() string {
	return awsutil.Prettify(s)
}

// Object containing configuration information for AWS Organizations.
type OrganizationConfiguration struct {
	_ struct{} `type:"structure"`

	// Flag to activate AWS Organization integration.
	//
	// EnableIntegration is a required field
	EnableIntegration *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s OrganizationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OrganizationConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OrganizationConfiguration"}

	if s.EnableIntegration == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnableIntegration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A set of attributes that describe a resource.
type ResourceInventory struct {
	_ struct{} `type:"structure"`

	// The platform of the resource.
	Platform *string `type:"string"`

	// Platform version of the resource in the inventory.
	PlatformVersion *string `type:"string"`

	// The ARN of the resource.
	ResourceArn *string `type:"string"`

	// Unique ID of the resource.
	ResourceId *string `type:"string"`

	// Unique ID of the account that owns the resource.
	ResourceOwningAccountId *string `type:"string"`

	// The type of resource.
	ResourceType enums.ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ResourceInventory) String() string {
	return awsutil.Prettify(s)
}

// Tag for a resource in a key-value format.
type Tag struct {
	_ struct{} `type:"structure"`

	// Key for the resource tag.
	Key *string `type:"string"`

	// Value for the resource tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}
