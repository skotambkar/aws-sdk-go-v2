// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetChangeTokenInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetChangeTokenInput) String() string {
	return awsutil.Prettify(s)
}

type GetChangeTokenOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used in the request. Use this value in a GetChangeTokenStatus
	// request to get the current status of the request.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetChangeTokenOutput) String() string {
	return awsutil.Prettify(s)
}
