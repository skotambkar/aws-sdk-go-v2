// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete an existing custom verification email template.
type DeleteCustomVerificationEmailTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the custom verification email template that you want to delete.
	//
	// TemplateName is a required field
	TemplateName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomVerificationEmailTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomVerificationEmailTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomVerificationEmailTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCustomVerificationEmailTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomVerificationEmailTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
