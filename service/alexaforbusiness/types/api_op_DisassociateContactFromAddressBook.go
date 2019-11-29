// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateContactFromAddressBookInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the address from which to disassociate the contact.
	//
	// AddressBookArn is a required field
	AddressBookArn *string `type:"string" required:"true"`

	// The ARN of the contact to disassociate from an address book.
	//
	// ContactArn is a required field
	ContactArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateContactFromAddressBookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateContactFromAddressBookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateContactFromAddressBookInput"}

	if s.AddressBookArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AddressBookArn"))
	}

	if s.ContactArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateContactFromAddressBookOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateContactFromAddressBookOutput) String() string {
	return awsutil.Prettify(s)
}