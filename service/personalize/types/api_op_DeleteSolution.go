// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSolutionInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the solution to delete.
	//
	// SolutionArn is a required field
	SolutionArn *string `locationName:"solutionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSolutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSolutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSolutionInput"}

	if s.SolutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSolutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSolutionOutput) String() string {
	return awsutil.Prettify(s)
}
