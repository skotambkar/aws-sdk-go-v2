// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteCoreDefinitionInput struct {
	_ struct{} `type:"structure"`

	// CoreDefinitionId is a required field
	CoreDefinitionId *string `location:"uri" locationName:"CoreDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCoreDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCoreDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCoreDefinitionInput"}

	if s.CoreDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CoreDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCoreDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCoreDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}