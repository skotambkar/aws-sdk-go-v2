// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLoggerDefinitionInput struct {
	_ struct{} `type:"structure"`

	// LoggerDefinitionId is a required field
	LoggerDefinitionId *string `location:"uri" locationName:"LoggerDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLoggerDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoggerDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLoggerDefinitionInput"}

	if s.LoggerDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggerDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLoggerDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLoggerDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}