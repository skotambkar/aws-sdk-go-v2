// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/enums"
)

// Represents the input of a CreateDeploymentConfig operation.
type CreateDeploymentConfigInput struct {
	_ struct{} `type:"structure"`

	// The destination platform type for the deployment (Lambda, Server, or ECS).
	ComputePlatform enums.ComputePlatform `locationName:"computePlatform" type:"string" enum:"true"`

	// The name of the deployment configuration to create.
	//
	// DeploymentConfigName is a required field
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string" required:"true"`

	// The minimum number of healthy instances that should be available at any time
	// during the deployment. There are two parameters expected in the input: type
	// and value.
	//
	// The type parameter takes either of the following values:
	//
	//    * HOST_COUNT: The value parameter represents the minimum number of healthy
	//    instances as an absolute value.
	//
	//    * FLEET_PERCENT: The value parameter represents the minimum number of
	//    healthy instances as a percentage of the total number of instances in
	//    the deployment. If you specify FLEET_PERCENT, at the start of the deployment,
	//    AWS CodeDeploy converts the percentage to the equivalent number of instance
	//    and rounds up fractional instances.
	//
	// The value parameter takes an integer.
	//
	// For example, to set a minimum of 95% healthy instance, specify a type of
	// FLEET_PERCENT and a value of 95.
	MinimumHealthyHosts *MinimumHealthyHosts `locationName:"minimumHealthyHosts" type:"structure"`

	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig *TrafficRoutingConfig `locationName:"trafficRoutingConfig" type:"structure"`
}

// String returns the string representation
func (s CreateDeploymentConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentConfigInput"}

	if s.DeploymentConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentConfigName"))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDeploymentConfig operation.
type CreateDeploymentConfigOutput struct {
	_ struct{} `type:"structure"`

	// A unique deployment configuration ID.
	DeploymentConfigId *string `locationName:"deploymentConfigId" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentConfigOutput) String() string {
	return awsutil.Prettify(s)
}
