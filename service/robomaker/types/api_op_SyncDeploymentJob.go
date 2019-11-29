// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/enums"
)

type SyncDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The target fleet for the synchronization.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SyncDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SyncDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SyncDeploymentJobInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SyncDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the synchronization request.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// Information about the deployment application configurations.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// Information about the deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The failure code if the job fails:
	//
	// InternalServiceError
	//
	// Internal service error.
	//
	// RobotApplicationCrash
	//
	// Robot application exited abnormally.
	//
	// SimulationApplicationCrash
	//
	// Simulation application exited abnormally.
	//
	// BadPermissionsRobotApplication
	//
	// Robot application bundle could not be downloaded.
	//
	// BadPermissionsSimulationApplication
	//
	// Simulation application bundle could not be downloaded.
	//
	// BadPermissionsS3Output
	//
	// Unable to publish outputs to customer-provided S3 bucket.
	//
	// BadPermissionsCloudwatchLogs
	//
	// Unable to publish logs to customer-provided CloudWatch Logs resource.
	//
	// SubnetIpLimitExceeded
	//
	// Subnet IP limit exceeded.
	//
	// ENILimitExceeded
	//
	// ENI limit exceeded.
	//
	// BadPermissionsUserCredentials
	//
	// Unable to use the Role provided.
	//
	// InvalidBundleRobotApplication
	//
	// Robot bundle cannot be extracted (invalid format, bundling error, or other
	// issue).
	//
	// InvalidBundleSimulationApplication
	//
	// Simulation bundle cannot be extracted (invalid format, bundling error, or
	// other issue).
	//
	// RobotApplicationVersionMismatchedEtag
	//
	// Etag for RobotApplication does not match value during version creation.
	//
	// SimulationApplicationVersionMismatchedEtag
	//
	// Etag for SimulationApplication does not match value during version creation.
	FailureCode enums.DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// The failure reason if the job fails.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// The status of the synchronization job.
	Status enums.DeploymentStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s SyncDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}