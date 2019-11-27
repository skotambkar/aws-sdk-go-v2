// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetServiceSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetServiceSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type GetServiceSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether cross-account discovery has been enabled.
	EnableCrossAccountsDiscovery *bool `type:"boolean"`

	// Indicates whether AWS Organizations has been integrated with License Manager
	// for cross-account discovery.
	OrganizationConfiguration *OrganizationConfiguration `type:"structure"`

	// Regional S3 bucket path for storing reports, license trail event data, discovery
	// data, etc.
	S3BucketArn *string `type:"string"`

	// SNS topic configured to receive notifications from License Manager.
	SnsTopicArn *string `type:"string"`
}

// String returns the string representation
func (s GetServiceSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
