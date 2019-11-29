// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/inspector/enums"
)

type DescribeFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the finding that you want to describe.
	//
	// FindingArns is a required field
	FindingArns []string `locationName:"findingArns" min:"1" type:"list" required:"true"`

	// The locale into which you want to translate a finding description, recommendation,
	// and the short description that identifies the finding.
	Locale enums.Locale `locationName:"locale" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFindingsInput"}

	if s.FindingArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingArns"))
	}
	if s.FindingArns != nil && len(s.FindingArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FindingArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFindingsOutput struct {
	_ struct{} `type:"structure"`

	// Finding details that cannot be described. An error code is provided for each
	// failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`

	// Information about the finding.
	//
	// Findings is a required field
	Findings []Finding `locationName:"findings" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeFindingsOutput) String() string {
	return awsutil.Prettify(s)
}
