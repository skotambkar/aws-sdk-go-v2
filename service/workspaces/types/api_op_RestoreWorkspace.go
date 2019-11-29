// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreWorkspaceInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreWorkspaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreWorkspaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreWorkspaceInput"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreWorkspaceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RestoreWorkspaceOutput) String() string {
	return awsutil.Prettify(s)
}
