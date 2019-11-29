// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/forecast/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
)

const opCreateForecast = "CreateForecast"

// CreateForecastRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Creates a forecast for each item in the TARGET_TIME_SERIES dataset that was
// used to train the predictor. This is known as inference. To retrieve the
// forecast for a single item at low latency, use the operation. To export the
// complete forecast into your Amazon Simple Storage Service (Amazon S3), use
// the CreateForecastExportJob operation.
//
// The range of the forecast is determined by the ForecastHorizon, specified
// in the CreatePredictor request, multiplied by the DataFrequency, specified
// in the CreateDataset request. When you query a forecast, you can request
// a specific date range within the complete forecast.
//
// To get a list of all your forecasts, use the ListForecasts operation.
//
// The forecasts generated by Amazon Forecast are in the same timezone as the
// dataset that was used to create the predictor.
//
// For more information, see howitworks-forecast.
//
// The Status of the forecast must be ACTIVE before you can query or export
// the forecast. Use the DescribeForecast operation to get the status.
//
//    // Example sending a request using CreateForecastRequest.
//    req := client.CreateForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreateForecast
func (c *Client) CreateForecastRequest(input *types.CreateForecastInput) CreateForecastRequest {
	op := &aws.Operation{
		Name:       opCreateForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateForecastInput{}
	}

	req := c.newRequest(op, input, &types.CreateForecastOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateForecastMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateForecastRequest{Request: req, Input: input, Copy: c.CreateForecastRequest}
}

// CreateForecastRequest is the request type for the
// CreateForecast API operation.
type CreateForecastRequest struct {
	*aws.Request
	Input *types.CreateForecastInput
	Copy  func(*types.CreateForecastInput) CreateForecastRequest
}

// Send marshals and sends the CreateForecast API request.
func (r CreateForecastRequest) Send(ctx context.Context) (*CreateForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateForecastResponse{
		CreateForecastOutput: r.Request.Data.(*types.CreateForecastOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateForecastResponse is the response type for the
// CreateForecast API operation.
type CreateForecastResponse struct {
	*types.CreateForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateForecast request.
func (r *CreateForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
