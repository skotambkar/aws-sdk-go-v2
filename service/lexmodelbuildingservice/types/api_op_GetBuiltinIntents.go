// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/enums"
)

type GetBuiltinIntentsInput struct {
	_ struct{} `type:"structure"`

	// A list of locales that the intent supports.
	Locale enums.Locale `location:"querystring" locationName:"locale" type:"string" enum:"true"`

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A pagination token that fetches the next page of intents. If this API call
	// is truncated, Amazon Lex returns a pagination token in the response. To fetch
	// the next page of intents, use the pagination token in the next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// Substring to match in built-in intent signatures. An intent will be returned
	// if any part of its signature matches the substring. For example, "xyz" matches
	// both "xyzabc" and "abcxyz." To find the signature for an intent, see Standard
	// Built-in Intents (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	SignatureContains *string `location:"querystring" locationName:"signatureContains" type:"string"`
}

// String returns the string representation
func (s GetBuiltinIntentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBuiltinIntentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBuiltinIntentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetBuiltinIntentsOutput struct {
	_ struct{} `type:"structure"`

	// An array of builtinIntentMetadata objects, one for each intent in the response.
	Intents []BuiltinIntentMetadata `locationName:"intents" type:"list"`

	// A pagination token that fetches the next page of intents. If the response
	// to this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token
	// in the next request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBuiltinIntentsOutput) String() string {
	return awsutil.Prettify(s)
}
