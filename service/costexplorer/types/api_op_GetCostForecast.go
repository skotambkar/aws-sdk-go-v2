// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/enums"
)

type GetCostForecastInput struct {
	_ struct{} `type:"structure"`

	// The filters that you want to use to filter your forecast. Cost Explorer API
	// supports all of the Cost Explorer filters.
	Filter *Expression `type:"structure"`

	// How granular you want the forecast to be. You can get 3 months of DAILY forecasts
	// or 12 months of MONTHLY forecasts.
	//
	// The GetCostForecast operation supports only DAILY and MONTHLY granularities.
	//
	// Granularity is a required field
	Granularity enums.Granularity `type:"string" required:"true" enum:"true"`

	// Which metric Cost Explorer uses to create your forecast. For more information
	// about blended and unblended rates, see Why does the "blended" annotation
	// appear on some line items in my bill? (https://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	//
	// Valid values for a GetCostForecast call are the following:
	//
	//    * AMORTIZED_COST
	//
	//    * BLENDED_COST
	//
	//    * NET_AMORTIZED_COST
	//
	//    * NET_UNBLENDED_COST
	//
	//    * UNBLENDED_COST
	//
	// Metric is a required field
	Metric enums.Metric `type:"string" required:"true" enum:"true"`

	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence
	// level. The higher the confidence level, the more confident Cost Explorer
	// is about the actual value falling in the prediction interval. Higher confidence
	// levels result in wider prediction intervals.
	PredictionIntervalLevel *int64 `min:"51" type:"integer"`

	// The period of time that you want the forecast to cover.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCostForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCostForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCostForecastInput"}
	if len(s.Granularity) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Granularity"))
	}
	if len(s.Metric) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if s.PredictionIntervalLevel != nil && *s.PredictionIntervalLevel < 51 {
		invalidParams.Add(aws.NewErrParamMinValue("PredictionIntervalLevel", 51))
	}

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

type GetCostForecastOutput struct {
	_ struct{} `type:"structure"`

	// The forecasts for your query, in order. For DAILY forecasts, this is a list
	// of days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []ForecastResult `type:"list"`

	// How much you are forecasted to spend over the forecast period, in USD.
	Total *MetricValue `type:"structure"`
}

// String returns the string representation
func (s GetCostForecastOutput) String() string {
	return awsutil.Prettify(s)
}
