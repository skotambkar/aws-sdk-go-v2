// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the applicable
	// IAM user or AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// The names of the deployment groups.
	//
	// DeploymentGroupNames is a required field
	DeploymentGroupNames []string `locationName:"deploymentGroupNames" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetDeploymentGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetDeploymentGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetDeploymentGroupsInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.DeploymentGroupNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentGroupNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deployment groups.
	DeploymentGroupsInfo []DeploymentGroupInfo `locationName:"deploymentGroupsInfo" type:"list"`

	// Information about errors that might have occurred during the API call.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`
}

// String returns the string representation
func (s BatchGetDeploymentGroupsOutput) String() string {
	return awsutil.Prettify(s)
}