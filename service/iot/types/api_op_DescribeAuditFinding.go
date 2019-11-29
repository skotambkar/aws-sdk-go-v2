// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAuditFindingInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a single audit finding. You can use this identifier
	// to apply mitigation actions to the finding.
	//
	// FindingId is a required field
	FindingId *string `location:"uri" locationName:"findingId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAuditFindingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAuditFindingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAuditFindingInput"}

	if s.FindingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingId"))
	}
	if s.FindingId != nil && len(*s.FindingId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FindingId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAuditFindingOutput struct {
	_ struct{} `type:"structure"`

	// The findings (results) of the audit.
	Finding *AuditFinding `locationName:"finding" type:"structure"`
}

// String returns the string representation
func (s DescribeAuditFindingOutput) String() string {
	return awsutil.Prettify(s)
}