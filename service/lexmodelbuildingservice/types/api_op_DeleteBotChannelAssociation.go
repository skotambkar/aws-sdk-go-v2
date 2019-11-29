// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteBotChannelAssociationInput struct {
	_ struct{} `type:"structure"`

	// An alias that points to the specific version of the Amazon Lex bot to which
	// this association is being made.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"aliasName" min:"1" type:"string" required:"true"`

	// The name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The name of the association. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBotChannelAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBotChannelAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBotChannelAssociationInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}
	if s.BotAlias != nil && len(*s.BotAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BotAlias", 1))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteBotChannelAssociationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBotChannelAssociationOutput) String() string {
	return awsutil.Prettify(s)
}