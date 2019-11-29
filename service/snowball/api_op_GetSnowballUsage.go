// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/snowball/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
)

const opGetSnowballUsage = "GetSnowballUsage"

// GetSnowballUsageRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns information about the Snowball service limit for your account, and
// also the number of Snowballs your account has in use.
//
// The default service limit for the number of Snowballs that you can have at
// one time is 1. If you want to increase your service limit, contact AWS Support.
//
//    // Example sending a request using GetSnowballUsageRequest.
//    req := client.GetSnowballUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetSnowballUsage
func (c *Client) GetSnowballUsageRequest(input *types.GetSnowballUsageInput) GetSnowballUsageRequest {
	op := &aws.Operation{
		Name:       opGetSnowballUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSnowballUsageInput{}
	}

	req := c.newRequest(op, input, &types.GetSnowballUsageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetSnowballUsageMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSnowballUsageRequest{Request: req, Input: input, Copy: c.GetSnowballUsageRequest}
}

// GetSnowballUsageRequest is the request type for the
// GetSnowballUsage API operation.
type GetSnowballUsageRequest struct {
	*aws.Request
	Input *types.GetSnowballUsageInput
	Copy  func(*types.GetSnowballUsageInput) GetSnowballUsageRequest
}

// Send marshals and sends the GetSnowballUsage API request.
func (r GetSnowballUsageRequest) Send(ctx context.Context) (*GetSnowballUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSnowballUsageResponse{
		GetSnowballUsageOutput: r.Request.Data.(*types.GetSnowballUsageOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSnowballUsageResponse is the response type for the
// GetSnowballUsage API operation.
type GetSnowballUsageResponse struct {
	*types.GetSnowballUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSnowballUsage request.
func (r *GetSnowballUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
