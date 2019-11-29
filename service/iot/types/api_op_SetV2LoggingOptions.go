// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type SetV2LoggingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The default logging level.
	DefaultLogLevel enums.LogLevel `locationName:"defaultLogLevel" type:"string" enum:"true"`

	// If true all logs are disabled. The default is false.
	DisableAllLogs *bool `locationName:"disableAllLogs" type:"boolean"`

	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn *string `locationName:"roleArn" type:"string"`
}

// String returns the string representation
func (s SetV2LoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

type SetV2LoggingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetV2LoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}