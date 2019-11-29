// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
)

const opGetBatchPrediction = "GetBatchPrediction"

// GetBatchPredictionRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a BatchPrediction that includes detailed metadata, status, and data
// file information for a Batch Prediction request.
//
//    // Example sending a request using GetBatchPredictionRequest.
//    req := client.GetBatchPredictionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetBatchPredictionRequest(input *types.GetBatchPredictionInput) GetBatchPredictionRequest {
	op := &aws.Operation{
		Name:       opGetBatchPrediction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetBatchPredictionInput{}
	}

	req := c.newRequest(op, input, &types.GetBatchPredictionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetBatchPredictionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetBatchPredictionRequest{Request: req, Input: input, Copy: c.GetBatchPredictionRequest}
}

// GetBatchPredictionRequest is the request type for the
// GetBatchPrediction API operation.
type GetBatchPredictionRequest struct {
	*aws.Request
	Input *types.GetBatchPredictionInput
	Copy  func(*types.GetBatchPredictionInput) GetBatchPredictionRequest
}

// Send marshals and sends the GetBatchPrediction API request.
func (r GetBatchPredictionRequest) Send(ctx context.Context) (*GetBatchPredictionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBatchPredictionResponse{
		GetBatchPredictionOutput: r.Request.Data.(*types.GetBatchPredictionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBatchPredictionResponse is the response type for the
// GetBatchPrediction API operation.
type GetBatchPredictionResponse struct {
	*types.GetBatchPredictionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBatchPrediction request.
func (r *GetBatchPredictionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
