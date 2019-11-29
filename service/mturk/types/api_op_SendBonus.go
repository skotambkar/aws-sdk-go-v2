// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendBonusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the assignment for which this bonus is paid.
	//
	// AssignmentId is a required field
	AssignmentId *string `min:"1" type:"string" required:"true"`

	// The Bonus amount is a US Dollar amount specified using a string (for example,
	// "5" represents $5.00 USD and "101.42" represents $101.42 USD). Do not include
	// currency symbols or currency codes.
	//
	// BonusAmount is a required field
	BonusAmount *string `type:"string" required:"true"`

	// A message that explains the reason for the bonus payment. The Worker receiving
	// the bonus can see this message.
	//
	// Reason is a required field
	Reason *string `type:"string" required:"true"`

	// A unique identifier for this request, which allows you to retry the call
	// on error without granting multiple bonuses. This is useful in cases such
	// as network timeouts where it is unclear whether or not the call succeeded
	// on the server. If the bonus already exists in the system from a previous
	// call using the same UniqueRequestToken, subsequent calls will return an error
	// with a message containing the request ID.
	UniqueRequestToken *string `min:"1" type:"string"`

	// The ID of the Worker being paid the bonus.
	//
	// WorkerId is a required field
	WorkerId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendBonusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendBonusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendBonusInput"}

	if s.AssignmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentId"))
	}
	if s.AssignmentId != nil && len(*s.AssignmentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentId", 1))
	}

	if s.BonusAmount == nil {
		invalidParams.Add(aws.NewErrParamRequired("BonusAmount"))
	}

	if s.Reason == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reason"))
	}
	if s.UniqueRequestToken != nil && len(*s.UniqueRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UniqueRequestToken", 1))
	}

	if s.WorkerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerId"))
	}
	if s.WorkerId != nil && len(*s.WorkerId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkerId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendBonusOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SendBonusOutput) String() string {
	return awsutil.Prettify(s)
}
