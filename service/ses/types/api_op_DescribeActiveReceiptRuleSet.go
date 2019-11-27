// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the metadata and receipt rules for the receipt
// rule set that is currently active. You use receipt rule sets to receive email
// with Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeActiveReceiptRuleSetInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeActiveReceiptRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the metadata and receipt rules for the receipt rule set that is
// currently active.
type DescribeActiveReceiptRuleSetOutput struct {
	_ struct{} `type:"structure"`

	// The metadata for the currently active receipt rule set. The metadata consists
	// of the rule set name and a timestamp of when the rule set was created.
	Metadata *ReceiptRuleSetMetadata `type:"structure"`

	// The receipt rules that belong to the active rule set.
	Rules []ReceiptRule `type:"list"`
}

// String returns the string representation
func (s DescribeActiveReceiptRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}
