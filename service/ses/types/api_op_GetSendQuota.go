// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSendQuotaInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSendQuotaInput) String() string {
	return awsutil.Prettify(s)
}

// Represents your Amazon SES daily sending quota, maximum send rate, and the
// number of emails you have sent in the last 24 hours.
type GetSendQuotaOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of emails the user is allowed to send in a 24-hour interval.
	// A value of -1 signifies an unlimited quota.
	Max24HourSend *float64 `type:"double"`

	// The maximum number of emails that Amazon SES can accept from the user's account
	// per second.
	//
	// The rate at which Amazon SES accepts the user's messages might be less than
	// the maximum send rate.
	MaxSendRate *float64 `type:"double"`

	// The number of emails sent during the previous 24 hours.
	SentLast24Hours *float64 `type:"double"`
}

// String returns the string representation
func (s GetSendQuotaOutput) String() string {
	return awsutil.Prettify(s)
}
