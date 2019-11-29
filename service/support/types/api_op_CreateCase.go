// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCaseInput struct {
	_ struct{} `type:"structure"`

	// The ID of a set of one or more attachments for the case. Create the set by
	// using AddAttachmentsToSet.
	AttachmentSetId *string `locationName:"attachmentSetId" type:"string"`

	// The category of problem for the AWS Support case.
	CategoryCode *string `locationName:"categoryCode" type:"string"`

	// A list of email addresses that AWS Support copies on case correspondence.
	CcEmailAddresses []string `locationName:"ccEmailAddresses" type:"list"`

	// The communication body text when you create an AWS Support case by calling
	// CreateCase.
	//
	// CommunicationBody is a required field
	CommunicationBody *string `locationName:"communicationBody" min:"1" type:"string" required:"true"`

	// The type of issue for the case. You can specify either "customer-service"
	// or "technical." If you do not indicate a value, the default is "technical."
	//
	// Service limit increases are not supported by the Support API; you must submit
	// service limit increase requests in Support Center (https://console.aws.amazon.com/support).
	IssueType *string `locationName:"issueType" type:"string"`

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters
	// must be passed explicitly for operations that take them.
	Language *string `locationName:"language" type:"string"`

	// The code for the AWS service returned by the call to DescribeServices.
	ServiceCode *string `locationName:"serviceCode" type:"string"`

	// The code for the severity level returned by the call to DescribeSeverityLevels.
	//
	// The availability of severity levels depends on the support plan for the account.
	SeverityCode *string `locationName:"severityCode" type:"string"`

	// The title of the AWS Support case.
	//
	// Subject is a required field
	Subject *string `locationName:"subject" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCaseInput"}

	if s.CommunicationBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommunicationBody"))
	}
	if s.CommunicationBody != nil && len(*s.CommunicationBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CommunicationBody", 1))
	}

	if s.Subject == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subject"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The AWS Support case ID returned by a successful completion of the CreateCase
// operation.
type CreateCaseOutput struct {
	_ struct{} `type:"structure"`

	// The AWS Support case ID requested or returned in the call. The case ID is
	// an alphanumeric string formatted as shown in this example: case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string `locationName:"caseId" type:"string"`
}

// String returns the string representation
func (s CreateCaseOutput) String() string {
	return awsutil.Prettify(s)
}