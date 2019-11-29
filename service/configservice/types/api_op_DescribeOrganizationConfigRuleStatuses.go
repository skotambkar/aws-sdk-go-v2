// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrganizationConfigRuleStatusesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of OrganizationConfigRuleStatuses returned on each page.
	// If you do no specify a number, AWS Config uses the default. The default is
	// 100.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// The names of organization config rules for which you want status details.
	// If you do not specify any names, AWS Config returns details for all your
	// organization AWS Confg rules.
	OrganizationConfigRuleNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRuleStatusesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeOrganizationConfigRuleStatusesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// A list of OrganizationConfigRuleStatus objects.
	OrganizationConfigRuleStatuses []OrganizationConfigRuleStatus `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRuleStatusesOutput) String() string {
	return awsutil.Prettify(s)
}