// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/enums"
)

type DescribeEnvironmentStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to get status information about.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `locationName:"environmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentStatusInput"}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEnvironmentStatusOutput struct {
	_ struct{} `type:"structure"`

	// Any informational message about the status of the environment.
	Message *string `locationName:"message" type:"string"`

	// The status of the environment. Available values include:
	//
	//    * connecting: The environment is connecting.
	//
	//    * creating: The environment is being created.
	//
	//    * deleting: The environment is being deleted.
	//
	//    * error: The environment is in an error state.
	//
	//    * ready: The environment is ready.
	//
	//    * stopped: The environment is stopped.
	//
	//    * stopping: The environment is stopping.
	Status enums.EnvironmentStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentStatusOutput) String() string {
	return awsutil.Prettify(s)
}
