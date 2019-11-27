// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMitigationActionInput struct {
	_ struct{} `type:"structure"`

	// The friendly name for the mitigation action. You can't change the name by
	// using UpdateMitigationAction. Instead, you must delete and re-create the
	// mitigation action with the new name.
	//
	// ActionName is a required field
	ActionName *string `location:"uri" locationName:"actionName" type:"string" required:"true"`

	// Defines the type of action and the parameters for that action.
	ActionParams *MitigationActionParams `locationName:"actionParams" type:"structure"`

	// The ARN of the IAM role that is used to apply the mitigation action.
	RoleArn *string `locationName:"roleArn" min:"20" type:"string"`
}

// String returns the string representation
func (s UpdateMitigationActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMitigationActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMitigationActionInput"}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.ActionParams != nil {
		if err := s.ActionParams.Validate(); err != nil {
			invalidParams.AddNested("ActionParams", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMitigationActionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN for the new mitigation action.
	ActionArn *string `locationName:"actionArn" type:"string"`

	// A unique identifier for the mitigation action.
	ActionId *string `locationName:"actionId" type:"string"`
}

// String returns the string representation
func (s UpdateMitigationActionOutput) String() string {
	return awsutil.Prettify(s)
}
