// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDeploymentTargetInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// The unique ID of a deployment target.
	TargetId *string `locationName:"targetId" type:"string"`
}

// String returns the string representation
func (s GetDeploymentTargetInput) String() string {
	return awsutil.Prettify(s)
}

type GetDeploymentTargetOutput struct {
	_ struct{} `type:"structure"`

	// A deployment target that contains information about a deployment such as
	// its status, lifecyle events, and when it was last updated. It also contains
	// metadata about the deployment target. The deployment target metadata depends
	// on the deployment target's type (instanceTarget, lambdaTarget, or ecsTarget).
	DeploymentTarget *DeploymentTarget `locationName:"deploymentTarget" type:"structure"`
}

// String returns the string representation
func (s GetDeploymentTargetOutput) String() string {
	return awsutil.Prettify(s)
}
