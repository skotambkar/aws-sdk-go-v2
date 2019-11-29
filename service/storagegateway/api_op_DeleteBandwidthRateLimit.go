// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDeleteBandwidthRateLimit = "DeleteBandwidthRateLimit"

// DeleteBandwidthRateLimitRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the bandwidth rate limits of a gateway. You can delete either the
// upload and download bandwidth rate limit, or you can delete both. If you
// delete only one of the limits, the other limit remains unchanged. To specify
// which gateway to work with, use the Amazon Resource Name (ARN) of the gateway
// in your request. This operation is supported for the stored volume, cached
// volume and tape gateway types.
//
//    // Example sending a request using DeleteBandwidthRateLimitRequest.
//    req := client.DeleteBandwidthRateLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteBandwidthRateLimit
func (c *Client) DeleteBandwidthRateLimitRequest(input *types.DeleteBandwidthRateLimitInput) DeleteBandwidthRateLimitRequest {
	op := &aws.Operation{
		Name:       opDeleteBandwidthRateLimit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteBandwidthRateLimitInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBandwidthRateLimitOutput{})
	return DeleteBandwidthRateLimitRequest{Request: req, Input: input, Copy: c.DeleteBandwidthRateLimitRequest}
}

// DeleteBandwidthRateLimitRequest is the request type for the
// DeleteBandwidthRateLimit API operation.
type DeleteBandwidthRateLimitRequest struct {
	*aws.Request
	Input *types.DeleteBandwidthRateLimitInput
	Copy  func(*types.DeleteBandwidthRateLimitInput) DeleteBandwidthRateLimitRequest
}

// Send marshals and sends the DeleteBandwidthRateLimit API request.
func (r DeleteBandwidthRateLimitRequest) Send(ctx context.Context) (*DeleteBandwidthRateLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBandwidthRateLimitResponse{
		DeleteBandwidthRateLimitOutput: r.Request.Data.(*types.DeleteBandwidthRateLimitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBandwidthRateLimitResponse is the response type for the
// DeleteBandwidthRateLimit API operation.
type DeleteBandwidthRateLimitResponse struct {
	*types.DeleteBandwidthRateLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBandwidthRateLimit request.
func (r *DeleteBandwidthRateLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
