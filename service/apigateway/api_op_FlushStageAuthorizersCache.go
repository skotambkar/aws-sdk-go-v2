// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opFlushStageAuthorizersCache = "FlushStageAuthorizersCache"

// FlushStageAuthorizersCacheRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Flushes all authorizer cache entries on a stage.
//
//    // Example sending a request using FlushStageAuthorizersCacheRequest.
//    req := client.FlushStageAuthorizersCacheRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) FlushStageAuthorizersCacheRequest(input *types.FlushStageAuthorizersCacheInput) FlushStageAuthorizersCacheRequest {
	op := &aws.Operation{
		Name:       opFlushStageAuthorizersCache,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}/cache/authorizers",
	}

	if input == nil {
		input = &types.FlushStageAuthorizersCacheInput{}
	}

	req := c.newRequest(op, input, &types.FlushStageAuthorizersCacheOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return FlushStageAuthorizersCacheRequest{Request: req, Input: input, Copy: c.FlushStageAuthorizersCacheRequest}
}

// FlushStageAuthorizersCacheRequest is the request type for the
// FlushStageAuthorizersCache API operation.
type FlushStageAuthorizersCacheRequest struct {
	*aws.Request
	Input *types.FlushStageAuthorizersCacheInput
	Copy  func(*types.FlushStageAuthorizersCacheInput) FlushStageAuthorizersCacheRequest
}

// Send marshals and sends the FlushStageAuthorizersCache API request.
func (r FlushStageAuthorizersCacheRequest) Send(ctx context.Context) (*FlushStageAuthorizersCacheResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &FlushStageAuthorizersCacheResponse{
		FlushStageAuthorizersCacheOutput: r.Request.Data.(*types.FlushStageAuthorizersCacheOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// FlushStageAuthorizersCacheResponse is the response type for the
// FlushStageAuthorizersCache API operation.
type FlushStageAuthorizersCacheResponse struct {
	*types.FlushStageAuthorizersCacheOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// FlushStageAuthorizersCache request.
func (r *FlushStageAuthorizersCacheResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
