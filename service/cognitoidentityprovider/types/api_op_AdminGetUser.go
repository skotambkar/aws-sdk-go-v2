// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/enums"
)

// Represents the request to get the specified user as an administrator.
type AdminGetUserInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID for the user pool where you want to get information about
	// the user.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name of the user you wish to retrieve.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminGetUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminGetUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminGetUserInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server from the request to get the specified
// user as an administrator.
type AdminGetUserOutput struct {
	_ struct{} `type:"structure"`

	// Indicates that the status is enabled.
	Enabled *bool `type:"boolean"`

	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about TOTP software
	// token MFA configurations. To look up information about either type of MFA
	// configuration, use the AdminGetUserResponse$UserMFASettingList response instead.
	MFAOptions []MFAOptionType `type:"list"`

	// The user's preferred MFA setting.
	PreferredMfaSetting *string `type:"string"`

	// An array of name-value pairs representing user attributes.
	UserAttributes []AttributeType `type:"list"`

	// The date the user was created.
	UserCreateDate *time.Time `type:"timestamp"`

	// The date the user was last modified.
	UserLastModifiedDate *time.Time `type:"timestamp"`

	// The MFA options that are enabled for the user. The possible values in this
	// list are SMS_MFA and SOFTWARE_TOKEN_MFA.
	UserMFASettingList []string `type:"list"`

	// The user status. Can be one of the following:
	//
	//    * UNCONFIRMED - User has been created but not confirmed.
	//
	//    * CONFIRMED - User has been confirmed.
	//
	//    * ARCHIVED - User is no longer active.
	//
	//    * COMPROMISED - User is disabled due to a potential security threat.
	//
	//    * UNKNOWN - User status is not known.
	//
	//    * RESET_REQUIRED - User is confirmed, but the user must request a code
	//    and reset his or her password before he or she can sign in.
	//
	//    * FORCE_CHANGE_PASSWORD - The user is confirmed and the user can sign
	//    in using a temporary password, but on first sign-in, the user must change
	//    his or her password to a new value before doing anything else.
	UserStatus enums.UserStatusType `type:"string" enum:"true"`

	// The user name of the user about whom you are receiving information.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminGetUserOutput) String() string {
	return awsutil.Prettify(s)
}
