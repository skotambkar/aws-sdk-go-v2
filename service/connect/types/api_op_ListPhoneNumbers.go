// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/connect/enums"
)

type ListPhoneNumbersInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The ISO country code.
	PhoneNumberCountryCodes []enums.PhoneNumberCountryCode `location:"querystring" locationName:"phoneNumberCountryCodes" type:"list"`

	// The type of phone number.
	PhoneNumberTypes []enums.PhoneNumberType `location:"querystring" locationName:"phoneNumberTypes" type:"list"`
}

// String returns the string representation
func (s ListPhoneNumbersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPhoneNumbersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPhoneNumbersInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListPhoneNumbersOutput struct {
	_ struct{} `type:"structure"`

	// If there are additional results, this is the token for the next set of results.
	NextToken *string `type:"string"`

	// Information about the phone numbers.
	PhoneNumberSummaryList []PhoneNumberSummary `type:"list"`
}

// String returns the string representation
func (s ListPhoneNumbersOutput) String() string {
	return awsutil.Prettify(s)
}
