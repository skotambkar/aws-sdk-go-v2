// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConformancePackComplianceInput struct {
	_ struct{} `type:"structure"`

	// Name of the conformance pack.
	//
	// ConformancePackName is a required field
	ConformancePackName *string `min:"1" type:"string" required:"true"`

	// A ConformancePackComplianceFilters object.
	Filters *ConformancePackComplianceFilters `type:"structure"`

	// The maximum number of AWS Config rules within a conformance pack are returned
	// on each page.
	Limit *int64 `type:"integer"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePackComplianceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConformancePackComplianceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConformancePackComplianceInput"}

	if s.ConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConformancePackName"))
	}
	if s.ConformancePackName != nil && len(*s.ConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConformancePackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeConformancePackComplianceOutput struct {
	_ struct{} `type:"structure"`

	// Name of the conformance pack.
	//
	// ConformancePackName is a required field
	ConformancePackName *string `min:"1" type:"string" required:"true"`

	// Returns a list of ConformancePackRuleCompliance objects.
	//
	// ConformancePackRuleComplianceList is a required field
	ConformancePackRuleComplianceList []ConformancePackRuleCompliance `type:"list" required:"true"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePackComplianceOutput) String() string {
	return awsutil.Prettify(s)
}