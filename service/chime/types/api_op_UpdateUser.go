// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/chime/enums"
)

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The user license type to update. This must be a supported license type for
	// the Amazon Chime account that the user belongs to.
	LicenseType enums.License `type:"string" enum:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`

	// The updated user details.
	User *User `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}
