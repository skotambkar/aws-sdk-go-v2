// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteIPSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The IPSetId of the IPSet that you want to delete. IPSetId is returned by
	// CreateIPSet and by ListIPSets.
	//
	// IPSetId is a required field
	IPSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIPSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.IPSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IPSetId"))
	}
	if s.IPSetId != nil && len(*s.IPSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IPSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIPSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteIPSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteIPSetOutput) String() string {
	return awsutil.Prettify(s)
}
