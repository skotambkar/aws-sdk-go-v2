// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteWebACLInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The WebACLId of the WebACL that you want to delete. WebACLId is returned
	// by CreateWebACL and by ListWebACLs.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWebACLInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteWebACLOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteWebACL request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteWebACLOutput) String() string {
	return awsutil.Prettify(s)
}
