// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type IsVpcPeeredInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s IsVpcPeeredInput) String() string {
	return awsutil.Prettify(s)
}

type IsVpcPeeredOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the Lightsail VPC is peered; otherwise, false.
	IsPeered *bool `locationName:"isPeered" type:"boolean"`
}

// String returns the string representation
func (s IsVpcPeeredOutput) String() string {
	return awsutil.Prettify(s)
}