// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Requests a list of AWS Cost and Usage reports owned by the account.
type DescribeReportDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results that AWS returns for the operation.
	MaxResults *int64 `min:"5" type:"integer"`

	// A generic string.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeReportDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReportDefinitionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReportDefinitionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// If the action is successful, the service sends back an HTTP 200 response.
type DescribeReportDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// A generic string.
	NextToken *string `type:"string"`

	// A list of AWS Cost and Usage reports owned by the account.
	ReportDefinitions []ReportDefinition `type:"list"`
}

// String returns the string representation
func (s DescribeReportDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}
