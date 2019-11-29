// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to list the IP address filters that exist under your
// AWS account. You use IP address filters when you receive email with Amazon
// SES. For more information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type ListReceiptFiltersInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListReceiptFiltersInput) String() string {
	return awsutil.Prettify(s)
}

// A list of IP address filters that exist under your AWS account.
type ListReceiptFiltersOutput struct {
	_ struct{} `type:"structure"`

	// A list of IP address filter data structures, which each consist of a name,
	// an IP address range, and whether to allow or block mail from it.
	Filters []ReceiptFilter `type:"list"`
}

// String returns the string representation
func (s ListReceiptFiltersOutput) String() string {
	return awsutil.Prettify(s)
}
