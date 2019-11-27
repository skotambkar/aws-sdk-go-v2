// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecastquery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery/types"
)

const opQueryForecast = "QueryForecast"

// QueryForecastRequest returns a request value for making API operation for
// Amazon Forecast Query Service.
//
// Retrieves a forecast filtered by the supplied criteria.
//
// The criteria is a key-value pair. The key is either item_id (or the equivalent
// non-timestamp, non-target field) from the TARGET_TIME_SERIES dataset, or
// one of the forecast dimensions specified as part of the FeaturizationConfig
// object.
//
// By default, the complete date range of the filtered forecast is returned.
// Optionally, you can request a specific date range within the forecast.
//
// The forecasts generated by Amazon Forecast are in the same timezone as the
// dataset that was used to create the predictor.
//
//    // Example sending a request using QueryForecastRequest.
//    req := client.QueryForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecastquery-2018-06-26/QueryForecast
func (c *Client) QueryForecastRequest(input *types.QueryForecastInput) QueryForecastRequest {
	op := &aws.Operation{
		Name:       opQueryForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.QueryForecastInput{}
	}

	req := c.newRequest(op, input, &types.QueryForecastOutput{})
	return QueryForecastRequest{Request: req, Input: input, Copy: c.QueryForecastRequest}
}

// QueryForecastRequest is the request type for the
// QueryForecast API operation.
type QueryForecastRequest struct {
	*aws.Request
	Input *types.QueryForecastInput
	Copy  func(*types.QueryForecastInput) QueryForecastRequest
}

// Send marshals and sends the QueryForecast API request.
func (r QueryForecastRequest) Send(ctx context.Context) (*QueryForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &QueryForecastResponse{
		QueryForecastOutput: r.Request.Data.(*types.QueryForecastOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// QueryForecastResponse is the response type for the
// QueryForecast API operation.
type QueryForecastResponse struct {
	*types.QueryForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// QueryForecast request.
func (r *QueryForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
