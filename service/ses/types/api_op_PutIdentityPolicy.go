// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to add or update a sending authorization policy for
// an identity. Sending authorization is an Amazon SES feature that enables
// you to authorize other senders to use your identities. For information, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type PutIdentityPolicyInput struct {
	_ struct{} `type:"structure"`

	// The identity that the policy will apply to. You can specify an identity by
	// using its name or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// To successfully call this API, you must own the identity.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// The text of the policy in JSON format. The policy cannot exceed 4 KB.
	//
	// For information about the syntax of sending authorization policies, see the
	// Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`

	// The name of the policy.
	//
	// The policy name cannot exceed 64 characters and can only include alphanumeric
	// characters, dashes, and underscores.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutIdentityPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutIdentityPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutIdentityPolicyInput"}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
	}

	if s.Policy == nil {
		invalidParams.Add(aws.NewErrParamRequired("Policy"))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type PutIdentityPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutIdentityPolicyOutput) String() string {
	return awsutil.Prettify(s)
}