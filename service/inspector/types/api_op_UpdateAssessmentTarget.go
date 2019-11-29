// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAssessmentTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment target that you want to update.
	//
	// AssessmentTargetArn is a required field
	AssessmentTargetArn *string `locationName:"assessmentTargetArn" min:"1" type:"string" required:"true"`

	// The name of the assessment target that you want to update.
	//
	// AssessmentTargetName is a required field
	AssessmentTargetName *string `locationName:"assessmentTargetName" min:"1" type:"string" required:"true"`

	// The ARN of the resource group that is used to specify the new resource group
	// to associate with the assessment target.
	ResourceGroupArn *string `locationName:"resourceGroupArn" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAssessmentTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAssessmentTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAssessmentTargetInput"}

	if s.AssessmentTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetArn"))
	}
	if s.AssessmentTargetArn != nil && len(*s.AssessmentTargetArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetArn", 1))
	}

	if s.AssessmentTargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetName"))
	}
	if s.AssessmentTargetName != nil && len(*s.AssessmentTargetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetName", 1))
	}
	if s.ResourceGroupArn != nil && len(*s.ResourceGroupArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAssessmentTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAssessmentTargetOutput) String() string {
	return awsutil.Prettify(s)
}