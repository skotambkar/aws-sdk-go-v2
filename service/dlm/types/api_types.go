// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dlm/enums"
)

// Specifies when to create snapshots of EBS volumes.
type CreateRule struct {
	_ struct{} `type:"structure"`

	// The interval between snapshots. The supported values are 2, 3, 4, 6, 8, 12,
	// and 24.
	//
	// Interval is a required field
	Interval *int64 `min:"1" type:"integer" required:"true"`

	// The interval unit.
	//
	// IntervalUnit is a required field
	IntervalUnit enums.IntervalUnitValues `type:"string" required:"true" enum:"true"`

	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time.
	Times []string `type:"list"`
}

// String returns the string representation
func (s CreateRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRule"}

	if s.Interval == nil {
		invalidParams.Add(aws.NewErrParamRequired("Interval"))
	}
	if s.Interval != nil && *s.Interval < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 1))
	}
	if len(s.IntervalUnit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IntervalUnit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies when to enable fast snapshot restore.
type FastRestoreRule struct {
	_ struct{} `type:"structure"`

	// The Availability Zones in which to enable fast snapshot restore.
	//
	// AvailabilityZones is a required field
	AvailabilityZones []string `min:"1" type:"list" required:"true"`

	// The number of snapshots to be enabled with fast snapshot restore.
	//
	// Count is a required field
	Count *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s FastRestoreRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FastRestoreRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FastRestoreRule"}

	if s.AvailabilityZones == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZones"))
	}
	if s.AvailabilityZones != nil && len(s.AvailabilityZones) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZones", 1))
	}

	if s.Count == nil {
		invalidParams.Add(aws.NewErrParamRequired("Count"))
	}
	if s.Count != nil && *s.Count < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Count", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Detailed information about a lifecycle policy.
type LifecyclePolicy struct {
	_ struct{} `type:"structure"`

	// The local date and time when the lifecycle policy was created.
	DateCreated *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The local date and time when the lifecycle policy was last modified.
	DateModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The description of the lifecycle policy.
	Description *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string `type:"string"`

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string `type:"string"`

	// The configuration of the lifecycle policy
	PolicyDetails *PolicyDetails `type:"structure"`

	// The identifier of the lifecycle policy.
	PolicyId *string `type:"string"`

	// The activation state of the lifecycle policy.
	State enums.GettablePolicyStateValues `type:"string" enum:"true"`

	// The description of the status.
	StatusMessage *string `type:"string"`

	// The tags.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s LifecyclePolicy) String() string {
	return awsutil.Prettify(s)
}

// Summary information about a lifecycle policy.
type LifecyclePolicySummary struct {
	_ struct{} `type:"structure"`

	// The description of the lifecycle policy.
	Description *string `type:"string"`

	// The identifier of the lifecycle policy.
	PolicyId *string `type:"string"`

	// The activation state of the lifecycle policy.
	State enums.GettablePolicyStateValues `type:"string" enum:"true"`

	// The tags.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s LifecyclePolicySummary) String() string {
	return awsutil.Prettify(s)
}

// Optional parameters that can be added to the policy. The set of valid parameters
// depends on the combination of policyType and resourceType values.
type Parameters struct {
	_ struct{} `type:"structure"`

	// When executing an EBS Snapshot Management – Instance policy, execute all
	// CreateSnapshots calls with the excludeBootVolume set to the supplied field.
	// Defaults to false. Only valid for EBS Snapshot Management – Instance policies.
	ExcludeBootVolume *bool `type:"boolean"`
}

// String returns the string representation
func (s Parameters) String() string {
	return awsutil.Prettify(s)
}

// Specifies the configuration of a lifecycle policy.
type PolicyDetails struct {
	_ struct{} `type:"structure"`

	// A set of optional parameters that can be provided by the policy.
	Parameters *Parameters `type:"structure"`

	// This field determines the valid target resource types and actions a policy
	// can manage. This field defaults to EBS_SNAPSHOT_MANAGEMENT if not present.
	PolicyType enums.PolicyTypeValues `type:"string" enum:"true"`

	// The resource type.
	ResourceTypes []enums.ResourceTypeValues `min:"1" type:"list"`

	// The schedule of policy-defined actions.
	Schedules []Schedule `min:"1" type:"list"`

	// The single tag that identifies targeted resources for this policy.
	TargetTags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s PolicyDetails) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PolicyDetails) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PolicyDetails"}
	if s.ResourceTypes != nil && len(s.ResourceTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceTypes", 1))
	}
	if s.Schedules != nil && len(s.Schedules) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Schedules", 1))
	}
	if s.TargetTags != nil && len(s.TargetTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetTags", 1))
	}
	if s.Schedules != nil {
		for i, v := range s.Schedules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Schedules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TargetTags != nil {
		for i, v := range s.TargetTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the number of snapshots to keep for each EBS volume.
type RetainRule struct {
	_ struct{} `type:"structure"`

	// The number of snapshots to keep for each volume, up to a maximum of 1000.
	//
	// Count is a required field
	Count *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s RetainRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetainRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetainRule"}

	if s.Count == nil {
		invalidParams.Add(aws.NewErrParamRequired("Count"))
	}
	if s.Count != nil && *s.Count < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Count", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies a schedule.
type Schedule struct {
	_ struct{} `type:"structure"`

	// Copy all user-defined tags on a source volume to snapshots of the volume
	// created by this policy.
	CopyTags *bool `type:"boolean"`

	// The create rule.
	CreateRule *CreateRule `type:"structure"`

	// Enable fast snapshot restore.
	FastRestoreRule *FastRestoreRule `type:"structure"`

	// The name of the schedule.
	Name *string `type:"string"`

	// The retain rule.
	RetainRule *RetainRule `type:"structure"`

	// The tags to apply to policy-created resources. These user-defined tags are
	// in addition to the AWS-added lifecycle tags.
	TagsToAdd []Tag `type:"list"`

	// A collection of key/value pairs with values determined dynamically when the
	// policy is executed. Keys may be any valid Amazon EC2 tag key. Values must
	// be in one of the two following formats: $(instance-id) or $(timestamp). Variable
	// tags are only valid for EBS Snapshot Management – Instance policies.
	VariableTags []Tag `type:"list"`
}

// String returns the string representation
func (s Schedule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Schedule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Schedule"}
	if s.CreateRule != nil {
		if err := s.CreateRule.Validate(); err != nil {
			invalidParams.AddNested("CreateRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.FastRestoreRule != nil {
		if err := s.FastRestoreRule.Validate(); err != nil {
			invalidParams.AddNested("FastRestoreRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.RetainRule != nil {
		if err := s.RetainRule.Validate(); err != nil {
			invalidParams.AddNested("RetainRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagsToAdd != nil {
		for i, v := range s.TagsToAdd {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagsToAdd", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VariableTags != nil {
		for i, v := range s.VariableTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "VariableTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies a tag for a resource.
type Tag struct {
	_ struct{} `type:"structure"`

	// The tag key.
	//
	// Key is a required field
	Key *string `type:"string" required:"true"`

	// The tag value.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}