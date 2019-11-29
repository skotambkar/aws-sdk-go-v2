// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutInput struct {
	_ struct{} `type:"structure"`

	// A NextToken string is used when you call the ListPhoneNumbersOptedOut action
	// to retrieve additional records that are available after the first page of
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPhoneNumbersOptedOutInput) String() string {
	return awsutil.Prettify(s)
}

// The response from the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutOutput struct {
	_ struct{} `type:"structure"`

	// A NextToken string is returned when you call the ListPhoneNumbersOptedOut
	// action if additional records are available after the first page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of phone numbers that are opted out of receiving SMS messages. The
	// list is paginated, and each page can contain up to 100 phone numbers.
	PhoneNumbers []string `locationName:"phoneNumbers" type:"list"`
}

// String returns the string representation
func (s ListPhoneNumbersOptedOutOutput) String() string {
	return awsutil.Prettify(s)
}
