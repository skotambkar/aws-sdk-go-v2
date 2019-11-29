// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to obtain information about the email-sending capabilities of your
// Amazon SES account.
type GetAccountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountInput) String() string {
	return awsutil.Prettify(s)
}

// A list of details about the email-sending capabilities of your Amazon SES
// account in the current AWS Region.
type GetAccountOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether or not the automatic warm-up feature is enabled for dedicated
	// IP addresses that are associated with your account.
	DedicatedIpAutoWarmupEnabled *bool `type:"boolean"`

	// The reputation status of your Amazon SES account. The status can be one of
	// the following:
	//
	//    * HEALTHY – There are no reputation-related issues that currently impact
	//    your account.
	//
	//    * PROBATION – We've identified potential issues with your Amazon SES
	//    account. We're placing your account under review while you work on correcting
	//    these issues.
	//
	//    * SHUTDOWN – Your account's ability to send email is currently paused
	//    because of an issue with the email sent from your account. When you correct
	//    the issue, you can contact us and request that your account's ability
	//    to send email is resumed.
	EnforcementStatus *string `type:"string"`

	// Indicates whether or not your account has production access in the current
	// AWS Region.
	//
	// If the value is false, then your account is in the sandbox. When your account
	// is in the sandbox, you can only send email to verified identities. Additionally,
	// the maximum number of emails you can send in a 24-hour period (your sending
	// quota) is 200, and the maximum number of emails you can send per second (your
	// maximum sending rate) is 1.
	//
	// If the value is true, then your account has production access. When your
	// account has production access, you can send email to any address. The sending
	// quota and maximum sending rate for your account vary based on your specific
	// use case.
	ProductionAccessEnabled *bool `type:"boolean"`

	// An object that contains information about the per-day and per-second sending
	// limits for your Amazon SES account in the current AWS Region.
	SendQuota *SendQuota `type:"structure"`

	// Indicates whether or not email sending is enabled for your Amazon SES account
	// in the current AWS Region.
	SendingEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s GetAccountOutput) String() string {
	return awsutil.Prettify(s)
}