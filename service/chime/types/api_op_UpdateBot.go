// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateBotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot ID.
	//
	// BotId is a required field
	BotId *string `location:"uri" locationName:"botId" type:"string" required:"true"`

	// When true, stops the specified bot from running in your account.
	Disabled *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBotInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.BotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateBotOutput struct {
	_ struct{} `type:"structure"`

	// The updated bot details.
	Bot *Bot `type:"structure"`
}

// String returns the string representation
func (s UpdateBotOutput) String() string {
	return awsutil.Prettify(s)
}
