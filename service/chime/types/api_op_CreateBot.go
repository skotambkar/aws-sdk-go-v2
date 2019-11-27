// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateBotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot display name.
	//
	// DisplayName is a required field
	DisplayName *string `type:"string" required:"true" sensitive:"true"`

	// The domain of the Amazon Chime Enterprise account.
	Domain *string `type:"string"`
}

// String returns the string representation
func (s CreateBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBotInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.DisplayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DisplayName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateBotOutput struct {
	_ struct{} `type:"structure"`

	// The bot details.
	Bot *Bot `type:"structure"`
}

// String returns the string representation
func (s CreateBotOutput) String() string {
	return awsutil.Prettify(s)
}
