// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RefreshTrustedAdvisorCheckInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the Trusted Advisor check to refresh. Note: Specifying
	// the check ID of a check that is automatically refreshed causes an InvalidParameterValue
	// error.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`
}

// String returns the string representation
func (s RefreshTrustedAdvisorCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RefreshTrustedAdvisorCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RefreshTrustedAdvisorCheckInput"}

	if s.CheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The current refresh status of a Trusted Advisor check.
type RefreshTrustedAdvisorCheckOutput struct {
	_ struct{} `type:"structure"`

	// The current refresh status for a check, including the amount of time until
	// the check is eligible for refresh.
	//
	// Status is a required field
	Status *TrustedAdvisorCheckRefreshStatus `locationName:"status" type:"structure" required:"true"`
}

// String returns the string representation
func (s RefreshTrustedAdvisorCheckOutput) String() string {
	return awsutil.Prettify(s)
}
