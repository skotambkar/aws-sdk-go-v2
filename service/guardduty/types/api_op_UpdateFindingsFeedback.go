// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/enums"
)

type UpdateFindingsFeedbackInput struct {
	_ struct{} `type:"structure"`

	// Additional feedback about the GuardDuty findings.
	Comments *string `locationName:"comments" type:"string"`

	// The ID of the detector associated with the findings to update feedback for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The feedback for the finding.
	//
	// Feedback is a required field
	Feedback enums.Feedback `locationName:"feedback" type:"string" required:"true" enum:"true"`

	// IDs of the findings that you want to mark as useful or not useful.
	//
	// FindingIds is a required field
	FindingIds []string `locationName:"findingIds" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateFindingsFeedbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFindingsFeedbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFindingsFeedbackInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if len(s.Feedback) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Feedback"))
	}

	if s.FindingIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateFindingsFeedbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateFindingsFeedbackOutput) String() string {
	return awsutil.Prettify(s)
}