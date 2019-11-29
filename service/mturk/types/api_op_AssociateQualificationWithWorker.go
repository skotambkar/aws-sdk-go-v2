// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateQualificationWithWorkerInput struct {
	_ struct{} `type:"structure"`

	// The value of the Qualification to assign.
	IntegerValue *int64 `type:"integer"`

	// The ID of the Qualification type to use for the assigned Qualification.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`

	// Specifies whether to send a notification email message to the Worker saying
	// that the qualification was assigned to the Worker. Note: this is true by
	// default.
	SendNotification *bool `type:"boolean"`

	// The ID of the Worker to whom the Qualification is being assigned. Worker
	// IDs are included with submitted HIT assignments and Qualification requests.
	//
	// WorkerId is a required field
	WorkerId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateQualificationWithWorkerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateQualificationWithWorkerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateQualificationWithWorkerInput"}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.QualificationTypeId != nil && len(*s.QualificationTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QualificationTypeId", 1))
	}

	if s.WorkerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerId"))
	}
	if s.WorkerId != nil && len(*s.WorkerId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkerId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateQualificationWithWorkerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateQualificationWithWorkerOutput) String() string {
	return awsutil.Prettify(s)
}