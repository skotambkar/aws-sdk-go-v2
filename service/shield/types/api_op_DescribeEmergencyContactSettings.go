// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEmergencyContactSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeEmergencyContactSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeEmergencyContactSettingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of email addresses that the DRT can use to contact you during a suspected
	// attack.
	EmergencyContactList []EmergencyContact `type:"list"`
}

// String returns the string representation
func (s DescribeEmergencyContactSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
