// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opGetAggregateResourceConfig = "GetAggregateResourceConfig"

// GetAggregateResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Returns configuration item that is aggregated for your specific resource
// in a specific source account and region.
//
//    // Example sending a request using GetAggregateResourceConfigRequest.
//    req := client.GetAggregateResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetAggregateResourceConfig
func (c *Client) GetAggregateResourceConfigRequest(input *types.GetAggregateResourceConfigInput) GetAggregateResourceConfigRequest {
	op := &aws.Operation{
		Name:       opGetAggregateResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetAggregateResourceConfigInput{}
	}

	req := c.newRequest(op, input, &types.GetAggregateResourceConfigOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetAggregateResourceConfigMarshaler{Input: input}.GetNamedBuildHandler())

	return GetAggregateResourceConfigRequest{Request: req, Input: input, Copy: c.GetAggregateResourceConfigRequest}
}

// GetAggregateResourceConfigRequest is the request type for the
// GetAggregateResourceConfig API operation.
type GetAggregateResourceConfigRequest struct {
	*aws.Request
	Input *types.GetAggregateResourceConfigInput
	Copy  func(*types.GetAggregateResourceConfigInput) GetAggregateResourceConfigRequest
}

// Send marshals and sends the GetAggregateResourceConfig API request.
func (r GetAggregateResourceConfigRequest) Send(ctx context.Context) (*GetAggregateResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAggregateResourceConfigResponse{
		GetAggregateResourceConfigOutput: r.Request.Data.(*types.GetAggregateResourceConfigOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAggregateResourceConfigResponse is the response type for the
// GetAggregateResourceConfig API operation.
type GetAggregateResourceConfigResponse struct {
	*types.GetAggregateResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAggregateResourceConfig request.
func (r *GetAggregateResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
