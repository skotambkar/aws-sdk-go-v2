// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/enums"
)

type DescribeDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDeploymentJobInput"}

	if s.Job == nil {
		invalidParams.Add(aws.NewErrParamRequired("Job"))
	}
	if s.Job != nil && len(*s.Job) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Job", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the deployment job was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The deployment application configuration.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The deployment job failure code.
	FailureCode enums.DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// A short description of the reason why the deployment job failed.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// A list of robot deployment summaries.
	RobotDeploymentSummary []RobotDeployment `locationName:"robotDeploymentSummary" type:"list"`

	// The status of the deployment job.
	Status enums.DeploymentStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the specified deployment job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}