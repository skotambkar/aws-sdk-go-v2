// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/configservice/enums"
)

type GetComplianceDetailsByConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// Filters the results by compliance.
	//
	// The allowed values are COMPLIANT, NON_COMPLIANT, and NOT_APPLICABLE.
	ComplianceTypes []enums.ComplianceType `type:"list"`

	// The name of the AWS Config rule for which you want compliance information.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// The maximum number of evaluation results returned on each page. The default
	// is 10. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetComplianceDetailsByConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetComplianceDetailsByConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetComplianceDetailsByConfigRuleInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetComplianceDetailsByConfigRuleOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the AWS resource complies with the specified AWS Config
	// rule.
	EvaluationResults []EvaluationResult `type:"list"`

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetComplianceDetailsByConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}
