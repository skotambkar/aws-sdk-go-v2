// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// The app ID. If you include this parameter, the command returns a description
	// of the commands associated with the specified app.
	AppId *string `type:"string"`

	// An array of deployment IDs to be described. If you include this parameter,
	// the command returns a description of the specified deployments. Otherwise,
	// it returns a description of every deployment.
	DeploymentIds []string `type:"list"`

	// The stack ID. If you include this parameter, the command returns a description
	// of the commands associated with the specified stack.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribeDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeDeployments request.
type DescribeDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// An array of Deployment objects that describe the deployments.
	Deployments []Deployment `type:"list"`
}

// String returns the string representation
func (s DescribeDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}
