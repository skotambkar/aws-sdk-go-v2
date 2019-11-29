// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccountSummaryInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful GetAccountSummary request.
type GetAccountSummaryOutput struct {
	_ struct{} `type:"structure"`

	// A set of key–value pairs containing information about IAM entity usage
	// and IAM quotas.
	SummaryMap map[string]int64 `type:"map"`
}

// String returns the string representation
func (s GetAccountSummaryOutput) String() string {
	return awsutil.Prettify(s)
}