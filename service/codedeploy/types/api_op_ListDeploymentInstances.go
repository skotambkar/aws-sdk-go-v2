// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/enums"
)

// Represents the input of a ListDeploymentInstances operation.
type ListDeploymentInstancesInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`

	// A subset of instances to list by status:
	//
	//    * Pending: Include those instances with pending deployments.
	//
	//    * InProgress: Include those instances where deployments are still in progress.
	//
	//    * Succeeded: Include those instances with successful deployments.
	//
	//    * Failed: Include those instances with failed deployments.
	//
	//    * Skipped: Include those instances with skipped deployments.
	//
	//    * Unknown: Include those instances with deployments in an unknown state.
	InstanceStatusFilter []enums.InstanceStatus `locationName:"instanceStatusFilter" type:"list"`

	// The set of instances in a blue/green deployment, either those in the original
	// environment ("BLUE") or those in the replacement environment ("GREEN"), for
	// which you want to view instance information.
	InstanceTypeFilter []enums.InstanceType `locationName:"instanceTypeFilter" type:"list"`

	// An identifier returned from the previous list deployment instances call.
	// It can be used to return the next set of deployment instances in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentInstancesInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListDeploymentInstances operation.
type ListDeploymentInstancesOutput struct {
	_ struct{} `type:"structure"`

	// A list of instance IDs.
	InstancesList []string `locationName:"instancesList" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment instances call to return the
	// next set of deployment instances in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentInstancesOutput) String() string {
	return awsutil.Prettify(s)
}