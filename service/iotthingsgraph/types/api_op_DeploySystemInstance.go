// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeploySystemInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the system instance. This value is returned by the CreateSystemInstance
	// action.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:deployment:DEPLOYMENTNAME
	Id *string `locationName:"id" type:"string"`
}

// String returns the string representation
func (s DeploySystemInstanceInput) String() string {
	return awsutil.Prettify(s)
}

type DeploySystemInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the Greengrass deployment used to deploy the system instance.
	GreengrassDeploymentId *string `locationName:"greengrassDeploymentId" type:"string"`

	// An object that contains summary information about a system instance that
	// was deployed.
	//
	// Summary is a required field
	Summary *SystemInstanceSummary `locationName:"summary" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeploySystemInstanceOutput) String() string {
	return awsutil.Prettify(s)
}