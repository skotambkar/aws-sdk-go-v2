// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccountSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type GetAccountSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Limits that are related to concurrency and code storage.
	AccountLimit *AccountLimit `type:"structure"`

	// The number of functions and amount of storage in use.
	AccountUsage *AccountUsage `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}