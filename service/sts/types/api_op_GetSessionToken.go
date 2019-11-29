// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSessionTokenInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, that the credentials should remain valid. Acceptable
	// durations for IAM user sessions range from 900 seconds (15 minutes) to 129,600
	// seconds (36 hours), with 43,200 seconds (12 hours) as the default. Sessions
	// for AWS account owners are restricted to a maximum of 3,600 seconds (one
	// hour). If the duration is longer than one hour, the session for AWS account
	// owners defaults to one hour.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// The identification number of the MFA device that is associated with the IAM
	// user who is making the GetSessionToken call. Specify this value if the IAM
	// user has a policy that requires MFA authentication. The value is either the
	// serial number for a hardware device (such as GAHT12345678) or an Amazon Resource
	// Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user).
	// You can find the device for an IAM user by going to the AWS Management Console
	// and viewing the user's security credentials.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@:/-
	SerialNumber *string `min:"9" type:"string"`

	// The value provided by the MFA device, if MFA is required. If any policy requires
	// the IAM user to submit an MFA code, specify this value. If MFA authentication
	// is required, the user must provide a code when requesting a set of temporary
	// security credentials. A user who fails to provide the code receives an "access
	// denied" response when requesting resources that require MFA authentication.
	//
	// The format for this parameter, as described by its regex pattern, is a sequence
	// of six numeric digits.
	TokenCode *string `min:"6" type:"string"`
}

// String returns the string representation
func (s GetSessionTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSessionTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSessionTokenInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(aws.NewErrParamMinLen("SerialNumber", 9))
	}
	if s.TokenCode != nil && len(*s.TokenCode) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("TokenCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetSessionToken request, including
// temporary AWS credentials that can be used to make AWS requests.
type GetSessionTokenOutput struct {
	_ struct{} `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// The size of the security token that STS API operations return is not fixed.
	// We strongly recommend that you make no assumptions about the maximum size.
	Credentials *Credentials `type:"structure"`
}

// String returns the string representation
func (s GetSessionTokenOutput) String() string {
	return awsutil.Prettify(s)
}