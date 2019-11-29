// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appstream/enums"
)

type UpdateStackInput struct {
	_ struct{} `type:"structure"`

	// The list of interface VPC endpoint (interface endpoint) objects. Users of
	// the stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []AccessEndpoint `min:"1" type:"list"`

	// The persistent application settings for users of a stack. When these settings
	// are enabled, changes that users make to applications and Windows settings
	// are automatically saved after each session and applied to the next session.
	ApplicationSettings *ApplicationSettings `type:"structure"`

	// The stack attributes to delete.
	AttributesToDelete []enums.StackAttribute `type:"list"`

	// Deletes the storage connectors currently enabled for the stack.
	DeleteStorageConnectors *bool `deprecated:"true" type:"boolean"`

	// The description to display.
	Description *string `type:"string"`

	// The stack name to display.
	DisplayName *string `type:"string"`

	// The domains where AppStream 2.0 streaming sessions can be embedded in an
	// iframe. You must approve the domains that you want to host embedded AppStream
	// 2.0 streaming sessions.
	EmbedHostDomains []string `min:"1" type:"list"`

	// The URL that users are redirected to after they choose the Send Feedback
	// link. If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string `type:"string"`

	// The name of the stack.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string `type:"string"`

	// The storage connectors to enable.
	StorageConnectors []StorageConnector `type:"list"`

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default, these actions are enabled.
	UserSettings []UserSetting `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStackInput"}
	if s.AccessEndpoints != nil && len(s.AccessEndpoints) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessEndpoints", 1))
	}
	if s.EmbedHostDomains != nil && len(s.EmbedHostDomains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EmbedHostDomains", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.UserSettings != nil && len(s.UserSettings) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserSettings", 1))
	}
	if s.AccessEndpoints != nil {
		for i, v := range s.AccessEndpoints {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccessEndpoints", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ApplicationSettings != nil {
		if err := s.ApplicationSettings.Validate(); err != nil {
			invalidParams.AddNested("ApplicationSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.StorageConnectors != nil {
		for i, v := range s.StorageConnectors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "StorageConnectors", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.UserSettings != nil {
		for i, v := range s.UserSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateStackOutput struct {
	_ struct{} `type:"structure"`

	// Information about the stack.
	Stack *Stack `type:"structure"`
}

// String returns the string representation
func (s UpdateStackOutput) String() string {
	return awsutil.Prettify(s)
}
