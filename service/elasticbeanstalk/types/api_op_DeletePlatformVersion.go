// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePlatformVersionInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the version of the custom platform.
	PlatformArn *string `type:"string"`
}

// String returns the string representation
func (s DeletePlatformVersionInput) String() string {
	return awsutil.Prettify(s)
}

type DeletePlatformVersionOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the version of the custom platform.
	PlatformSummary *PlatformSummary `type:"structure"`
}

// String returns the string representation
func (s DeletePlatformVersionOutput) String() string {
	return awsutil.Prettify(s)
}
