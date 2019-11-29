// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/enums"
)

type DescribeRobotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the robot to be described.
	//
	// Robot is a required field
	Robot *string `locationName:"robot" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRobotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRobotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRobotInput"}

	if s.Robot == nil {
		invalidParams.Add(aws.NewErrParamRequired("Robot"))
	}
	if s.Robot != nil && len(*s.Robot) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Robot", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRobotOutput struct {
	_ struct{} `type:"structure"`

	// The target architecture of the robot application.
	Architecture enums.Architecture `locationName:"architecture" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string `locationName:"fleetArn" min:"1" type:"string"`

	// The Greengrass group id.
	GreengrassGroupId *string `locationName:"greengrassGroupId" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string `locationName:"lastDeploymentJob" min:"1" type:"string"`

	// The time of the last deployment job.
	LastDeploymentTime *time.Time `locationName:"lastDeploymentTime" type:"timestamp"`

	// The name of the robot.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The status of the fleet.
	Status enums.RobotStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the specified robot.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeRobotOutput) String() string {
	return awsutil.Prettify(s)
}