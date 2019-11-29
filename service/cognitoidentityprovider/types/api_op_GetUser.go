// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get information about the user.
type GetUserInput struct {
	_ struct{} `type:"structure"`

	// The access token returned by the server response to get information about
	// the user.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s GetUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server from the request to get information
// about the user.
type GetUserOutput struct {
	_ struct{} `type:"structure"`

	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about TOTP software
	// token MFA configurations. To look up information about either type of MFA
	// configuration, use the use the GetUserResponse$UserMFASettingList response
	// instead.
	MFAOptions []MFAOptionType `type:"list"`

	// The user's preferred MFA setting.
	PreferredMfaSetting *string `type:"string"`

	// An array of name-value pairs representing user attributes.
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// UserAttributes is a required field
	UserAttributes []AttributeType `type:"list" required:"true"`

	// The MFA options that are enabled for the user. The possible values in this
	// list are SMS_MFA and SOFTWARE_TOKEN_MFA.
	UserMFASettingList []string `type:"list"`

	// The user name of the user you wish to retrieve from the get user request.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s GetUserOutput) String() string {
	return awsutil.Prettify(s)
}
