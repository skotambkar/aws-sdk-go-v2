// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateApiKeyInput struct {
	_ struct{} `type:"structure"`

	// The ID for your GraphQL API.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A description of the purpose of the API key.
	Description *string `locationName:"description" type:"string"`

	// The time from creation time after which the API key expires. The date is
	// represented as seconds since the epoch, rounded down to the nearest hour.
	// The default value for this parameter is 7 days from creation time. For more
	// information, see .
	Expires *int64 `locationName:"expires" type:"long"`
}

// String returns the string representation
func (s CreateApiKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApiKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApiKeyInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateApiKeyOutput struct {
	_ struct{} `type:"structure"`

	// The API key.
	ApiKey *ApiKey `locationName:"apiKey" type:"structure"`
}

// String returns the string representation
func (s CreateApiKeyOutput) String() string {
	return awsutil.Prettify(s)
}