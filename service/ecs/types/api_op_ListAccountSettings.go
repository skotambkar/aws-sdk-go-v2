// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type ListAccountSettingsInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to return the effective settings. If true, the account
	// settings for the root user or the default setting for the principalArn are
	// returned. If false, the account settings for the principalArn are returned
	// if they are set. Otherwise, no account settings are returned.
	EffectiveSettings *bool `locationName:"effectiveSettings" type:"boolean"`

	// The maximum number of account setting results returned by ListAccountSettings
	// in paginated output. When this parameter is used, ListAccountSettings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListAccountSettings request with the returned nextToken value. This
	// value can be between 1 and 10. If this parameter is not used, then ListAccountSettings
	// returns up to 10 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The resource name you want to list the account settings for.
	Name enums.SettingName `locationName:"name" type:"string" enum:"true"`

	// The nextToken value returned from a ListAccountSettings request indicating
	// that more results are available to fulfill the request and further calls
	// will be needed. If maxResults was provided, it is possible the number of
	// results to be fewer than maxResults.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ARN of the principal, which can be an IAM user, IAM role, or the root
	// user. If this field is omitted, the account settings are listed only for
	// the authenticated user.
	PrincipalArn *string `locationName:"principalArn" type:"string"`

	// The value of the account settings with which to filter results. You must
	// also specify an account setting name to use this parameter.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s ListAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type ListAccountSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListAccountSettings request. When
	// the results of a ListAccountSettings request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The account settings for the resource.
	Settings []Setting `locationName:"settings" type:"list"`
}

// String returns the string representation
func (s ListAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
