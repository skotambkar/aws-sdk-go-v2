// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAssessmentTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment template that you want to delete.
	//
	// AssessmentTemplateArn is a required field
	AssessmentTemplateArn *string `locationName:"assessmentTemplateArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssessmentTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssessmentTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssessmentTemplateInput"}

	if s.AssessmentTemplateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTemplateArn"))
	}
	if s.AssessmentTemplateArn != nil && len(*s.AssessmentTemplateArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTemplateArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAssessmentTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssessmentTemplateOutput) String() string {
	return awsutil.Prettify(s)
}