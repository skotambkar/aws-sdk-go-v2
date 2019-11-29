// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConfigRuleEvaluationStatusInput struct {
	_ struct{} `type:"structure"`

	// The name of the AWS managed Config rules for which you want status information.
	// If you do not specify any names, AWS Config returns status information for
	// all AWS managed Config rules that you use.
	ConfigRuleNames []string `type:"list"`

	// The number of rule evaluation results that you want returned.
	//
	// This parameter is required if the rule limit for your account is more than
	// the default of 150 rules.
	//
	// For information about requesting a rule limit increase, see AWS Config Limits
	// (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config)
	// in the AWS General Reference Guide.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConfigRuleEvaluationStatusInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeConfigRuleEvaluationStatusOutput struct {
	_ struct{} `type:"structure"`

	// Status information about your AWS managed Config rules.
	ConfigRulesEvaluationStatus []ConfigRuleEvaluationStatus `type:"list"`

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConfigRuleEvaluationStatusOutput) String() string {
	return awsutil.Prettify(s)
}
