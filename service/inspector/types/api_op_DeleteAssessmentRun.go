// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAssessmentRunInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment run that you want to delete.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssessmentRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssessmentRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssessmentRunInput"}

	if s.AssessmentRunArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentRunArn"))
	}
	if s.AssessmentRunArn != nil && len(*s.AssessmentRunArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAssessmentRunOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssessmentRunOutput) String() string {
	return awsutil.Prettify(s)
}
