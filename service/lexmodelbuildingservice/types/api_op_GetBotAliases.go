// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBotAliasesInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The maximum number of aliases to return in the response. The default is 50. .
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in bot alias names. An alias will be returned if any part
	// of its name matches the substring. For example, "xyz" matches both "xyzabc"
	// and "abcxyz."
	NameContains *string `location:"querystring" locationName:"nameContains" min:"1" type:"string"`

	// A pagination token for fetching the next page of aliases. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotAliasesInput"}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetBotAliasesOutput struct {
	_ struct{} `type:"structure"`

	// An array of BotAliasMetadata objects, each describing a bot alias.
	BotAliases []BotAliasMetadata `type:"list"`

	// A pagination token for fetching next page of aliases. If the response to
	// this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotAliasesOutput) String() string {
	return awsutil.Prettify(s)
}