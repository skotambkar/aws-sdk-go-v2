// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf/enums"
)

type GetChangeTokenStatusInput struct {
	_ struct{} `type:"structure"`

	// The change token for which you want to get the status. This change token
	// was previously returned in the GetChangeToken response.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetChangeTokenStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetChangeTokenStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetChangeTokenStatusInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetChangeTokenStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the change token.
	ChangeTokenStatus enums.ChangeTokenStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetChangeTokenStatusOutput) String() string {
	return awsutil.Prettify(s)
}