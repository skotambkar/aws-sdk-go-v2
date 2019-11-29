// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListLoggerDefinitionVersionsInput struct {
	_ struct{} `type:"structure"`

	// LoggerDefinitionId is a required field
	LoggerDefinitionId *string `location:"uri" locationName:"LoggerDefinitionId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListLoggerDefinitionVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLoggerDefinitionVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLoggerDefinitionVersionsInput"}

	if s.LoggerDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggerDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListLoggerDefinitionVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	Versions []VersionInformation `type:"list"`
}

// String returns the string representation
func (s ListLoggerDefinitionVersionsOutput) String() string {
	return awsutil.Prettify(s)
}