// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAutoScalingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// One or more Availability Zones for the group.
	AvailabilityZones []string `min:"1" type:"list"`

	// The amount of time, in seconds, after a scaling activity completes before
	// another scaling activity can start. The default value is 300. This cooldown
	// period is not used when a scaling-specific cooldown is specified.
	//
	// Cooldown periods are not supported for target tracking scaling policies,
	// step scaling policies, or scheduled scaling. For more information, see Scaling
	// Cooldowns (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	DefaultCooldown *int64 `type:"integer"`

	// The number of EC2 instances that should be running in the Auto Scaling group.
	// This number must be greater than or equal to the minimum size of the group
	// and less than or equal to the maximum size of the group.
	DesiredCapacity *int64 `type:"integer"`

	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before
	// checking the health status of an EC2 instance that has come into service.
	// The default value is 0.
	//
	// For more information, see Health Check Grace Period (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// Conditional: This parameter is required if you are adding an ELB health check.
	HealthCheckGracePeriod *int64 `type:"integer"`

	// The service to use for the health checks. The valid values are EC2 and ELB.
	// If you configure an Auto Scaling group to use ELB health checks, it considers
	// the instance unhealthy if it fails either the EC2 status checks or the load
	// balancer health checks.
	HealthCheckType *string `min:"1" type:"string"`

	// The name of the launch configuration. If you specify LaunchConfigurationName
	// in your update request, you can't specify LaunchTemplate or MixedInstancesPolicy.
	LaunchConfigurationName *string `min:"1" type:"string"`

	// The launch template and version to use to specify the updates. If you specify
	// LaunchTemplate in your update request, you can't specify LaunchConfigurationName
	// or MixedInstancesPolicy.
	//
	// For more information, see LaunchTemplateSpecification (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchTemplateSpecification.html)
	// in the Amazon EC2 Auto Scaling API Reference.
	LaunchTemplate *LaunchTemplateSpecification `type:"structure"`

	// The maximum amount of time, in seconds, that an instance can be in service.
	//
	// Valid Range: Minimum value of 604800.
	MaxInstanceLifetime *int64 `type:"integer"`

	// The maximum size of the Auto Scaling group.
	MaxSize *int64 `type:"integer"`

	// The minimum size of the Auto Scaling group.
	MinSize *int64 `type:"integer"`

	// An embedded object that specifies a mixed instances policy.
	//
	// In your call to UpdateAutoScalingGroup, you can make changes to the policy
	// that is specified. All optional parameters are left unchanged if not specified.
	//
	// For more information, see MixedInstancesPolicy (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_MixedInstancesPolicy.html)
	// in the Amazon EC2 Auto Scaling API Reference and Auto Scaling Groups with
	// Multiple Instance Types and Purchase Options (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MixedInstancesPolicy *MixedInstancesPolicy `type:"structure"`

	// Indicates whether newly launched instances are protected from termination
	// by Amazon EC2 Auto Scaling when scaling in.
	//
	// For more information about preventing instances from terminating on scale
	// in, see Instance Protection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection)
	// in the Amazon EC2 Auto Scaling User Guide.
	NewInstancesProtectedFromScaleIn *bool `type:"boolean"`

	// The name of the placement group into which to launch your instances, if any.
	// A placement group is a logical grouping of instances within a single Availability
	// Zone. You cannot specify multiple Availability Zones and a placement group.
	// For more information, see Placement Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	PlacementGroup *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling
	// group uses to call other AWS services on your behalf. For more information,
	// see Service-Linked Roles (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	ServiceLinkedRoleARN *string `min:"1" type:"string"`

	// A standalone termination policy or a list of termination policies used to
	// select the instance to terminate. The policies are executed in the order
	// that they are listed.
	//
	// For more information, see Controlling Which Instances Auto Scaling Terminates
	// During Scale In (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TerminationPolicies []string `type:"list"`

	// A comma-separated list of subnet IDs for virtual private cloud (VPC).
	//
	// If you specify VPCZoneIdentifier with AvailabilityZones, the subnets that
	// you specify for this parameter must reside in those Availability Zones.
	VPCZoneIdentifier *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAutoScalingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAutoScalingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAutoScalingGroupInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}
	if s.AvailabilityZones != nil && len(s.AvailabilityZones) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZones", 1))
	}
	if s.HealthCheckType != nil && len(*s.HealthCheckType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HealthCheckType", 1))
	}
	if s.LaunchConfigurationName != nil && len(*s.LaunchConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchConfigurationName", 1))
	}
	if s.PlacementGroup != nil && len(*s.PlacementGroup) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementGroup", 1))
	}
	if s.ServiceLinkedRoleARN != nil && len(*s.ServiceLinkedRoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceLinkedRoleARN", 1))
	}
	if s.VPCZoneIdentifier != nil && len(*s.VPCZoneIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VPCZoneIdentifier", 1))
	}
	if s.LaunchTemplate != nil {
		if err := s.LaunchTemplate.Validate(); err != nil {
			invalidParams.AddNested("LaunchTemplate", err.(aws.ErrInvalidParams))
		}
	}
	if s.MixedInstancesPolicy != nil {
		if err := s.MixedInstancesPolicy.Validate(); err != nil {
			invalidParams.AddNested("MixedInstancesPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAutoScalingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAutoScalingGroupOutput) String() string {
	return awsutil.Prettify(s)
}