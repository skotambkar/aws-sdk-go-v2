// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for an EstimateTemplateCost action.
type EstimateTemplateCostInput struct {
	_ struct{} `type:"structure"`

	// A list of Parameter structures that specify input parameters.
	Parameters []Parameter `type:"list"`

	// Structure containing the template body with a minimum length of 1 byte and
	// a maximum length of 51,200 bytes. (For more information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.)
	//
	// Conditional: You must pass TemplateBody or TemplateURL. If both are passed,
	// only TemplateBody is used.
	TemplateBody *string `min:"1" type:"string"`

	// Location of file containing the template body. The URL must point to a template
	// that is located in an Amazon S3 bucket. For more information, go to Template
	// Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must pass TemplateURL or TemplateBody. If both are passed,
	// only TemplateBody is used.
	TemplateURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EstimateTemplateCostInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EstimateTemplateCostInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EstimateTemplateCostInput"}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateURL != nil && len(*s.TemplateURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateURL", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a EstimateTemplateCost action.
type EstimateTemplateCostOutput struct {
	_ struct{} `type:"structure"`

	// An AWS Simple Monthly Calculator URL with a query string that describes the
	// resources required to run the template.
	Url *string `type:"string"`
}

// String returns the string representation
func (s EstimateTemplateCostOutput) String() string {
	return awsutil.Prettify(s)
}