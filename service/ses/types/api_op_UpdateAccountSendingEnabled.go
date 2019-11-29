// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to enable or disable the email sending capabilities
// for your entire Amazon SES account.
type UpdateAccountSendingEnabledInput struct {
	_ struct{} `type:"structure"`

	// Describes whether email sending is enabled or disabled for your Amazon SES
	// account in the current AWS Region.
	Enabled *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateAccountSendingEnabledInput) String() string {
	return awsutil.Prettify(s)
}

type UpdateAccountSendingEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccountSendingEnabledOutput) String() string {
	return awsutil.Prettify(s)
}