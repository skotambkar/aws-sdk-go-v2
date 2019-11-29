// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type SetV2LoggingLevelInput struct {
	_ struct{} `type:"structure"`

	// The log level.
	//
	// LogLevel is a required field
	LogLevel enums.LogLevel `locationName:"logLevel" type:"string" required:"true" enum:"true"`

	// The log target.
	//
	// LogTarget is a required field
	LogTarget *LogTarget `locationName:"logTarget" type:"structure" required:"true"`
}

// String returns the string representation
func (s SetV2LoggingLevelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetV2LoggingLevelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetV2LoggingLevelInput"}
	if len(s.LogLevel) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LogLevel"))
	}

	if s.LogTarget == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogTarget"))
	}
	if s.LogTarget != nil {
		if err := s.LogTarget.Validate(); err != nil {
			invalidParams.AddNested("LogTarget", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetV2LoggingLevelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetV2LoggingLevelOutput) String() string {
	return awsutil.Prettify(s)
}