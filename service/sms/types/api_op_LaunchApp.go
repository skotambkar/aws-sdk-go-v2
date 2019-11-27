// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type LaunchAppInput struct {
	_ struct{} `type:"structure"`

	// ID of the application to launch.
	AppId *string `locationName:"appId" type:"string"`
}

// String returns the string representation
func (s LaunchAppInput) String() string {
	return awsutil.Prettify(s)
}

type LaunchAppOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s LaunchAppOutput) String() string {
	return awsutil.Prettify(s)
}
