// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRetentionConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the retention configuration to delete.
	//
	// RetentionConfigurationName is a required field
	RetentionConfigurationName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRetentionConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRetentionConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRetentionConfigurationInput"}

	if s.RetentionConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetentionConfigurationName"))
	}
	if s.RetentionConfigurationName != nil && len(*s.RetentionConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RetentionConfigurationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRetentionConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRetentionConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
