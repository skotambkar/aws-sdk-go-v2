// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConferencePreferenceInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetConferencePreferenceInput) String() string {
	return awsutil.Prettify(s)
}

type GetConferencePreferenceOutput struct {
	_ struct{} `type:"structure"`

	// The conference preference.
	Preference *ConferencePreference `type:"structure"`
}

// String returns the string representation
func (s GetConferencePreferenceOutput) String() string {
	return awsutil.Prettify(s)
}
