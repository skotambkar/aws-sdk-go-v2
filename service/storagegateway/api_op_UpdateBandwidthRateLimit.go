// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opUpdateBandwidthRateLimit = "UpdateBandwidthRateLimit"

// UpdateBandwidthRateLimitRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you
// don't set a bandwidth rate limit, the existing rate limit remains. This operation
// is supported for the stored volume, cached volume and tape gateway types.'
//
// By default, a gateway's bandwidth rate limits are not set. If you don't set
// any limit, the gateway does not have any limitations on its bandwidth usage
// and could potentially use the maximum available bandwidth.
//
// To specify which gateway to update, use the Amazon Resource Name (ARN) of
// the gateway in your request.
//
//    // Example sending a request using UpdateBandwidthRateLimitRequest.
//    req := client.UpdateBandwidthRateLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateBandwidthRateLimit
func (c *Client) UpdateBandwidthRateLimitRequest(input *types.UpdateBandwidthRateLimitInput) UpdateBandwidthRateLimitRequest {
	op := &aws.Operation{
		Name:       opUpdateBandwidthRateLimit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateBandwidthRateLimitInput{}
	}

	req := c.newRequest(op, input, &types.UpdateBandwidthRateLimitOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateBandwidthRateLimitMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateBandwidthRateLimitRequest{Request: req, Input: input, Copy: c.UpdateBandwidthRateLimitRequest}
}

// UpdateBandwidthRateLimitRequest is the request type for the
// UpdateBandwidthRateLimit API operation.
type UpdateBandwidthRateLimitRequest struct {
	*aws.Request
	Input *types.UpdateBandwidthRateLimitInput
	Copy  func(*types.UpdateBandwidthRateLimitInput) UpdateBandwidthRateLimitRequest
}

// Send marshals and sends the UpdateBandwidthRateLimit API request.
func (r UpdateBandwidthRateLimitRequest) Send(ctx context.Context) (*UpdateBandwidthRateLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBandwidthRateLimitResponse{
		UpdateBandwidthRateLimitOutput: r.Request.Data.(*types.UpdateBandwidthRateLimitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBandwidthRateLimitResponse is the response type for the
// UpdateBandwidthRateLimit API operation.
type UpdateBandwidthRateLimitResponse struct {
	*types.UpdateBandwidthRateLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBandwidthRateLimit request.
func (r *UpdateBandwidthRateLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
