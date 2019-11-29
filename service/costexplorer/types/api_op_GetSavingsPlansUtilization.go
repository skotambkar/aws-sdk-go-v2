// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/enums"
)

type GetSavingsPlansUtilizationInput struct {
	_ struct{} `type:"structure"`

	// Filters Savings Plans utilization coverage data for active Savings Plans
	// dimensions. You can filter data with the following dimensions:
	//
	//    * LINKED_ACCOUNT
	//
	//    * SAVINGS_PLAN_ARN
	//
	//    * SAVINGS_PLANS_TYPE
	//
	//    * REGION
	//
	//    * PAYMENT_OPTION
	//
	//    * INSTANCE_TYPE_FAMILY
	//
	// GetSavingsPlansUtilization uses the same Expression (http://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	Filter *Expression `type:"structure"`

	// The granularity of the Amazon Web Services utillization data for your Savings
	// Plans.
	//
	// The GetSavingsPlansUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity enums.Granularity `type:"string" enum:"true"`

	// The time period that you want the usage and costs for. The Start date must
	// be within 13 months. The End date must be after the Start date, and before
	// the current date. Future dates can't be used as an End date.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansUtilizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSavingsPlansUtilizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSavingsPlansUtilizationInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSavingsPlansUtilizationOutput struct {
	_ struct{} `type:"structure"`

	// The amount of cost/commitment you used your Savings Plans. This allows you
	// to specify date ranges.
	SavingsPlansUtilizationsByTime []SavingsPlansUtilizationByTime `type:"list"`

	// The total amount of cost/commitment that you used your Savings Plans, regardless
	// of date ranges.
	//
	// Total is a required field
	Total *SavingsPlansUtilizationAggregates `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansUtilizationOutput) String() string {
	return awsutil.Prettify(s)
}