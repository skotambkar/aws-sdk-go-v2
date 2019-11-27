// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCoreDefinitionInput struct {
	_ struct{} `type:"structure"`

	// CoreDefinitionId is a required field
	CoreDefinitionId *string `location:"uri" locationName:"CoreDefinitionId" type:"string" required:"true"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateCoreDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCoreDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCoreDefinitionInput"}

	if s.CoreDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CoreDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCoreDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCoreDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}
