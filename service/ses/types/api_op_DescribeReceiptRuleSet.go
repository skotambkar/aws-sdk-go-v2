// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the details of a receipt rule set. You use
// receipt rule sets to receive email with Amazon SES. For more information,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeReceiptRuleSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the receipt rule set to describe.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeReceiptRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReceiptRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReceiptRuleSetInput"}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the details of the specified receipt rule set.
type DescribeReceiptRuleSetOutput struct {
	_ struct{} `type:"structure"`

	// The metadata for the receipt rule set, which consists of the rule set name
	// and the timestamp of when the rule set was created.
	Metadata *ReceiptRuleSetMetadata `type:"structure"`

	// A list of the receipt rules that belong to the specified receipt rule set.
	Rules []ReceiptRule `type:"list"`
}

// String returns the string representation
func (s DescribeReceiptRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}