// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to list the configuration sets associated with your
// AWS account. Configuration sets enable you to publish email sending events.
// For information about using configuration sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsInput struct {
	_ struct{} `type:"structure"`

	// The number of configuration sets to return.
	MaxItems *int64 `type:"integer"`

	// A token returned from a previous call to ListConfigurationSets to indicate
	// the position of the configuration set in the configuration set list.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsInput) String() string {
	return awsutil.Prettify(s)
}

// A list of configuration sets associated with your AWS account. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsOutput struct {
	_ struct{} `type:"structure"`

	// A list of configuration sets.
	ConfigurationSets []ConfigurationSet `type:"list"`

	// A token indicating that there are additional configuration sets available
	// to be listed. Pass this token to successive calls of ListConfigurationSets.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsOutput) String() string {
	return awsutil.Prettify(s)
}