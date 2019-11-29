// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestartAppServerInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to restart the server for.
	//
	// Condition: You must specify either this or an EnvironmentName, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentId *string `type:"string"`

	// The name of the environment to restart the server for.
	//
	// Condition: You must specify either this or an EnvironmentId, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s RestartAppServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestartAppServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestartAppServerInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestartAppServerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RestartAppServerOutput) String() string {
	return awsutil.Prettify(s)
}