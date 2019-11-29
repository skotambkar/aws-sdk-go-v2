// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Creates a Cost and Usage Report.
type PutReportDefinitionInput struct {
	_ struct{} `type:"structure"`

	// Represents the output of the PutReportDefinition operation. The content consists
	// of the detailed metadata and data file information.
	//
	// ReportDefinition is a required field
	ReportDefinition *ReportDefinition `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutReportDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutReportDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutReportDefinitionInput"}

	if s.ReportDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReportDefinition"))
	}
	if s.ReportDefinition != nil {
		if err := s.ReportDefinition.Validate(); err != nil {
			invalidParams.AddNested("ReportDefinition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// If the action is successful, the service sends back an HTTP 200 response
// with an empty HTTP body.
type PutReportDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutReportDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}
