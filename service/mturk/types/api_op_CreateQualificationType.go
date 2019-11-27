// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mturk/enums"
)

type CreateQualificationTypeInput struct {
	_ struct{} `type:"structure"`

	// The answers to the Qualification test specified in the Test parameter, in
	// the form of an AnswerKey data structure.
	//
	// Constraints: Must not be longer than 65535 bytes.
	//
	// Constraints: None. If not specified, you must process Qualification requests
	// manually.
	AnswerKey *string `type:"string"`

	// Specifies whether requests for the Qualification type are granted immediately,
	// without prompting the Worker with a Qualification test.
	//
	// Constraints: If the Test parameter is specified, this parameter cannot be
	// true.
	AutoGranted *bool `type:"boolean"`

	// The Qualification value to use for automatically granted Qualifications.
	// This parameter is used only if the AutoGranted parameter is true.
	AutoGrantedValue *int64 `type:"integer"`

	// A long description for the Qualification type. On the Amazon Mechanical Turk
	// website, the long description is displayed when a Worker examines a Qualification
	// type.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// One or more words or phrases that describe the Qualification type, separated
	// by commas. The keywords of a type make the type easier to find during a search.
	Keywords *string `type:"string"`

	// The name you give to the Qualification type. The type name is used to represent
	// the Qualification to Workers, and to find the type using a Qualification
	// type search. It must be unique across all of your Qualification types.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The initial status of the Qualification type.
	//
	// Constraints: Valid values are: Active | Inactive
	//
	// QualificationTypeStatus is a required field
	QualificationTypeStatus enums.QualificationTypeStatus `type:"string" required:"true" enum:"true"`

	// The number of seconds that a Worker must wait after requesting a Qualification
	// of the Qualification type before the worker can retry the Qualification request.
	//
	// Constraints: None. If not specified, retries are disabled and Workers can
	// request a Qualification of this type only once, even if the Worker has not
	// been granted the Qualification. It is not possible to disable retries for
	// a Qualification type after it has been created with retries enabled. If you
	// want to disable retries, you must delete existing retry-enabled Qualification
	// type and then create a new Qualification type with retries disabled.
	RetryDelayInSeconds *int64 `type:"long"`

	// The questions for the Qualification test a Worker must answer correctly to
	// obtain a Qualification of this type. If this parameter is specified, TestDurationInSeconds
	// must also be specified.
	//
	// Constraints: Must not be longer than 65535 bytes. Must be a QuestionForm
	// data structure. This parameter cannot be specified if AutoGranted is true.
	//
	// Constraints: None. If not specified, the Worker may request the Qualification
	// without answering any questions.
	Test *string `type:"string"`

	// The number of seconds the Worker has to complete the Qualification test,
	// starting from the time the Worker requests the Qualification.
	TestDurationInSeconds *int64 `type:"long"`
}

// String returns the string representation
func (s CreateQualificationTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateQualificationTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateQualificationTypeInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.QualificationTypeStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateQualificationTypeOutput struct {
	_ struct{} `type:"structure"`

	// The created Qualification type, returned as a QualificationType data structure.
	QualificationType *QualificationType `type:"structure"`
}

// String returns the string representation
func (s CreateQualificationTypeOutput) String() string {
	return awsutil.Prettify(s)
}
