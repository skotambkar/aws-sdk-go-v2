// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RejectQualificationRequestInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Qualification request, as returned by the ListQualificationRequests
	// operation.
	//
	// QualificationRequestId is a required field
	QualificationRequestId *string `type:"string" required:"true"`

	// A text message explaining why the request was rejected, to be shown to the
	// Worker who made the request.
	Reason *string `type:"string"`
}

// String returns the string representation
func (s RejectQualificationRequestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectQualificationRequestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectQualificationRequestInput"}

	if s.QualificationRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RejectQualificationRequestOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RejectQualificationRequestOutput) String() string {
	return awsutil.Prettify(s)
}
