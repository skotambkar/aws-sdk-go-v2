// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete an email template. For more information, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type DeleteTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the template to be deleted.
	//
	// TemplateName is a required field
	TemplateName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
