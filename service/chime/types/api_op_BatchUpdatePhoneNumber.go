// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchUpdatePhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The request containing the phone number IDs and product types or calling
	// names to update.
	//
	// UpdatePhoneNumberRequestItems is a required field
	UpdatePhoneNumberRequestItems []UpdatePhoneNumberRequestItem `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchUpdatePhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUpdatePhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUpdatePhoneNumberInput"}

	if s.UpdatePhoneNumberRequestItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdatePhoneNumberRequestItems"))
	}
	if s.UpdatePhoneNumberRequestItems != nil {
		for i, v := range s.UpdatePhoneNumberRequestItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UpdatePhoneNumberRequestItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchUpdatePhoneNumberOutput struct {
	_ struct{} `type:"structure"`

	// If the action fails for one or more of the phone numbers in the request,
	// a list of the phone numbers is returned, along with error codes and error
	// messages.
	PhoneNumberErrors []PhoneNumberError `type:"list"`
}

// String returns the string representation
func (s BatchUpdatePhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}
