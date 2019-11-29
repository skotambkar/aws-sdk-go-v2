// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/enums"
)

// Request to list an environment's upcoming and in-progress managed actions.
type DescribeEnvironmentManagedActionsInput struct {
	_ struct{} `type:"structure"`

	// The environment ID of the target environment.
	EnvironmentId *string `type:"string"`

	// The name of the target environment.
	EnvironmentName *string `type:"string"`

	// To show only actions with a particular status, specify a status.
	Status enums.ActionStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentManagedActionsInput) String() string {
	return awsutil.Prettify(s)
}

// The result message containing a list of managed actions.
type DescribeEnvironmentManagedActionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of upcoming and in-progress managed actions.
	ManagedActions []ManagedAction `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeEnvironmentManagedActionsOutput) String() string {
	return awsutil.Prettify(s)
}
