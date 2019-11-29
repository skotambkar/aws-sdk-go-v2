// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the update project operation.
type UpdateProjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project whose name you wish to update.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// The number of minutes a test run in the project will execute before it times
	// out.
	DefaultJobTimeoutMinutes *int64 `locationName:"defaultJobTimeoutMinutes" type:"integer"`

	// A string representing the new name of the project that you are updating.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProjectInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of an update project request.
type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`

	// The project you wish to update.
	Project *Project `locationName:"project" type:"structure"`
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return awsutil.Prettify(s)
}
