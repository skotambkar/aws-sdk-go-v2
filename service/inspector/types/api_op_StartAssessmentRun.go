// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartAssessmentRunInput struct {
	_ struct{} `type:"structure"`

	// You can specify the name for the assessment run. The name must be unique
	// for the assessment template whose ARN is used to start the assessment run.
	AssessmentRunName *string `locationName:"assessmentRunName" min:"1" type:"string"`

	// The ARN of the assessment template of the assessment run that you want to
	// start.
	//
	// AssessmentTemplateArn is a required field
	AssessmentTemplateArn *string `locationName:"assessmentTemplateArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAssessmentRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAssessmentRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAssessmentRunInput"}
	if s.AssessmentRunName != nil && len(*s.AssessmentRunName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunName", 1))
	}

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

type StartAssessmentRunOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment run that has been started.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAssessmentRunOutput) String() string {
	return awsutil.Prettify(s)
}
