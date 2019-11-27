// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetObjectTorrent = "GetObjectTorrent"

// GetObjectTorrentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Return torrent files from a bucket.
//
//    // Example sending a request using GetObjectTorrentRequest.
//    req := client.GetObjectTorrentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTorrent
func (c *Client) GetObjectTorrentRequest(input *types.GetObjectTorrentInput) GetObjectTorrentRequest {
	op := &aws.Operation{
		Name:       opGetObjectTorrent,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?torrent",
	}

	if input == nil {
		input = &types.GetObjectTorrentInput{}
	}

	req := c.newRequest(op, input, &types.GetObjectTorrentOutput{})
	return GetObjectTorrentRequest{Request: req, Input: input, Copy: c.GetObjectTorrentRequest}
}

// GetObjectTorrentRequest is the request type for the
// GetObjectTorrent API operation.
type GetObjectTorrentRequest struct {
	*aws.Request
	Input *types.GetObjectTorrentInput
	Copy  func(*types.GetObjectTorrentInput) GetObjectTorrentRequest
}

// Send marshals and sends the GetObjectTorrent API request.
func (r GetObjectTorrentRequest) Send(ctx context.Context) (*GetObjectTorrentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectTorrentResponse{
		GetObjectTorrentOutput: r.Request.Data.(*types.GetObjectTorrentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectTorrentResponse is the response type for the
// GetObjectTorrent API operation.
type GetObjectTorrentResponse struct {
	*types.GetObjectTorrentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectTorrent request.
func (r *GetObjectTorrentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
