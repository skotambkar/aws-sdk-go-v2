// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to delete.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `locationName:"environmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEnvironmentInput"}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEnvironmentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}
