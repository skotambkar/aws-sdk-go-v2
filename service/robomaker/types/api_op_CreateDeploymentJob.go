// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/enums"
)

type CreateDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The deployment application configuration.
	//
	// DeploymentApplicationConfigs is a required field
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list" required:"true"`

	// The requested deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet to deploy.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`

	// A map that contains tag keys and tag values that are attached to the deployment
	// job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentJobInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DeploymentApplicationConfigs == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentApplicationConfigs"))
	}
	if s.DeploymentApplicationConfigs != nil && len(s.DeploymentApplicationConfigs) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentApplicationConfigs", 1))
	}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}
	if s.DeploymentApplicationConfigs != nil {
		for i, v := range s.DeploymentApplicationConfigs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DeploymentApplicationConfigs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.DeploymentConfig != nil {
		if err := s.DeploymentConfig.Validate(); err != nil {
			invalidParams.AddNested("DeploymentConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The deployment application configuration.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The failure code of the simulation job if it failed:
	//
	// BadPermissionError
	//
	// AWS Greengrass requires a service-level role permission to access other services.
	// The role must include the AWSGreengrassResourceAccessRolePolicy managed policy
	// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSGreengrassResourceAccessRolePolicy$jsonEditor).
	//
	// ExtractingBundleFailure
	//
	// The robot application could not be extracted from the bundle.
	//
	// FailureThresholdBreached
	//
	// The percentage of robots that could not be updated exceeded the percentage
	// set for the deployment.
	//
	// GreengrassDeploymentFailed
	//
	// The robot application could not be deployed to the robot.
	//
	// GreengrassGroupVersionDoesNotExist
	//
	// The AWS Greengrass group or version associated with a robot is missing.
	//
	// InternalServerError
	//
	// An internal error has occurred. Retry your request, but if the problem persists,
	// contact us with details.
	//
	// MissingRobotApplicationArchitecture
	//
	// The robot application does not have a source that matches the architecture
	// of the robot.
	//
	// MissingRobotDeploymentResource
	//
	// One or more of the resources specified for the robot application are missing.
	// For example, does the robot application have the correct launch package and
	// launch file?
	//
	// PostLaunchFileFailure
	//
	// The post-launch script failed.
	//
	// PreLaunchFileFailure
	//
	// The pre-launch script failed.
	//
	// ResourceNotFound
	//
	// One or more deployment resources are missing. For example, do robot application
	// source bundles still exist?
	//
	// RobotDeploymentNoResponse
	//
	// There is no response from the robot. It might not be powered on or connected
	// to the internet.
	FailureCode enums.DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// The failure reason of the deployment job if it failed.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The target fleet for the deployment job.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// The status of the deployment job.
	Status enums.DeploymentStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the deployment job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}
