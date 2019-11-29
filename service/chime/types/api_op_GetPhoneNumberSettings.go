// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPhoneNumberSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetPhoneNumberSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type GetPhoneNumberSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The default outbound calling name for the account.
	CallingName *string `type:"string" sensitive:"true"`

	// The updated outbound calling name timestamp, in ISO 8601 format.
	CallingNameUpdatedTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetPhoneNumberSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
