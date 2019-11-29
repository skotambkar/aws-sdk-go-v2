// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SkipWaitTimeForInstanceTerminationInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a blue/green deployment for which you want to skip the instance
	// termination wait time.
	DeploymentId *string `locationName:"deploymentId" type:"string"`
}

// String returns the string representation
func (s SkipWaitTimeForInstanceTerminationInput) String() string {
	return awsutil.Prettify(s)
}

type SkipWaitTimeForInstanceTerminationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SkipWaitTimeForInstanceTerminationOutput) String() string {
	return awsutil.Prettify(s)
}