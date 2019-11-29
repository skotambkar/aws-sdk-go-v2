// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteHsmConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Redshift HSM configuration to be deleted.
	//
	// HsmConfigurationIdentifier is a required field
	HsmConfigurationIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHsmConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHsmConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHsmConfigurationInput"}

	if s.HsmConfigurationIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmConfigurationIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteHsmConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHsmConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}