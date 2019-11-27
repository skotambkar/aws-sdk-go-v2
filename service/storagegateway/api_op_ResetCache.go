// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opResetCache = "ResetCache"

// ResetCacheRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Resets all cache disks that have encountered a error and makes the disks
// available for reconfiguration as cache storage. If your cache disk encounters
// a error, the gateway prevents read and write operations on virtual tapes
// in the gateway. For example, an error can occur when a disk is corrupted
// or removed from the gateway. When a cache is reset, the gateway loses its
// cache storage. At this point you can reconfigure the disks as cache disks.
// This operation is only supported in the cached volume and tape types.
//
// If the cache disk you are resetting contains data that has not been uploaded
// to Amazon S3 yet, that data can be lost. After you reset cache disks, there
// will be no configured cache disks left in the gateway, so you must configure
// at least one new cache disk for your gateway to function properly.
//
//    // Example sending a request using ResetCacheRequest.
//    req := client.ResetCacheRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ResetCache
func (c *Client) ResetCacheRequest(input *types.ResetCacheInput) ResetCacheRequest {
	op := &aws.Operation{
		Name:       opResetCache,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResetCacheInput{}
	}

	req := c.newRequest(op, input, &types.ResetCacheOutput{})
	return ResetCacheRequest{Request: req, Input: input, Copy: c.ResetCacheRequest}
}

// ResetCacheRequest is the request type for the
// ResetCache API operation.
type ResetCacheRequest struct {
	*aws.Request
	Input *types.ResetCacheInput
	Copy  func(*types.ResetCacheInput) ResetCacheRequest
}

// Send marshals and sends the ResetCache API request.
func (r ResetCacheRequest) Send(ctx context.Context) (*ResetCacheResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetCacheResponse{
		ResetCacheOutput: r.Request.Data.(*types.ResetCacheOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetCacheResponse is the response type for the
// ResetCache API operation.
type ResetCacheResponse struct {
	*types.ResetCacheOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetCache request.
func (r *ResetCacheResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
