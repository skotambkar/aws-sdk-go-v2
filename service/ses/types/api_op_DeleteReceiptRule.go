// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DeleteReceiptRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the receipt rule to delete.
	//
	// RuleName is a required field
	RuleName *string `type:"string" required:"true"`

	// The name of the receipt rule set that contains the receipt rule to delete.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReceiptRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReceiptRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReceiptRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type DeleteReceiptRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteReceiptRuleOutput) String() string {
	return awsutil.Prettify(s)
}