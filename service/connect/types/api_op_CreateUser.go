// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the user account in the directory used for identity management.
	// If Amazon Connect cannot access the directory, you can specify this identifier
	// to authenticate users. If you include the identifier, we assume that Amazon
	// Connect cannot access the directory. Otherwise, the identity information
	// is used to authenticate users from your directory.
	//
	// This parameter is required if you are using an existing directory for identity
	// management in Amazon Connect when Amazon Connect cannot access your directory
	// to authenticate users. If you are using SAML for identity management and
	// include this parameter, an error is returned.
	DirectoryUserId *string `type:"string"`

	// The identifier of the hierarchy group for the user.
	HierarchyGroupId *string `type:"string"`

	// The information about the identity of the user.
	IdentityInfo *UserIdentityInfo `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The password for the user account. A password is required if you are using
	// Amazon Connect for identity management. Otherwise, it is an error to include
	// a password.
	Password *string `type:"string"`

	// The phone settings for the user.
	//
	// PhoneConfig is a required field
	PhoneConfig *UserPhoneConfig `type:"structure" required:"true"`

	// The identifier of the routing profile for the user.
	//
	// RoutingProfileId is a required field
	RoutingProfileId *string `type:"string" required:"true"`

	// The identifier of the security profile for the user.
	//
	// SecurityProfileIds is a required field
	SecurityProfileIds []string `min:"1" type:"list" required:"true"`

	// One or more tags.
	Tags map[string]string `min:"1" type:"map"`

	// The user name for the account. For instances not using SAML for identity
	// management, the user name can include up to 20 characters. If you are using
	// SAML for identity management, the user name can include up to 64 characters
	// from [a-zA-Z0-9_-.\@]+.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.PhoneConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneConfig"))
	}

	if s.RoutingProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoutingProfileId"))
	}

	if s.SecurityProfileIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileIds"))
	}
	if s.SecurityProfileIds != nil && len(s.SecurityProfileIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileIds", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}
	if s.IdentityInfo != nil {
		if err := s.IdentityInfo.Validate(); err != nil {
			invalidParams.AddNested("IdentityInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.PhoneConfig != nil {
		if err := s.PhoneConfig.Validate(); err != nil {
			invalidParams.AddNested("PhoneConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user account.
	UserArn *string `type:"string"`

	// The identifier of the user account.
	UserId *string `type:"string"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}
