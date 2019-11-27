// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePolicyInput struct {
	_ struct{} `type:"structure"`

	// A friendly description of the policy.
	//
	// Typically used to store information about the permissions defined in the
	// policy. For example, "Grants access to production DynamoDB tables."
	//
	// The policy description is immutable. After a value is assigned, it cannot
	// be changed.
	Description *string `type:"string"`

	// The path for the policy.
	//
	// For more information about paths, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`

	// The JSON policy document that you want to use as the content for the new
	// policy.
	//
	// You must provide policies in JSON format in IAM. However, for AWS CloudFormation
	// templates formatted in YAML, you can provide the policy in JSON or YAML format.
	// AWS CloudFormation always converts a YAML policy to JSON format before submitting
	// it to IAM.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// PolicyDocument is a required field
	PolicyDocument *string `min:"1" type:"string" required:"true"`

	// The friendly name of the policy.
	//
	// IAM user, group, role, and policy names must be unique within the account.
	// Names are not distinguished by case. For example, you cannot create resources
	// named both "MyResource" and "myresource".
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePolicyInput"}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if s.PolicyDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyDocument"))
	}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyDocument", 1))
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

// Contains the response to a successful CreatePolicy request.
type CreatePolicyOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new policy.
	Policy *Policy `type:"structure"`
}

// String returns the string representation
func (s CreatePolicyOutput) String() string {
	return awsutil.Prettify(s)
}
