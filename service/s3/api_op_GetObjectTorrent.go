// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetObjectTorrent = "GetObjectTorrent"

// GetObjectTorrentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Return torrent files from a bucket. BitTorrent can save you bandwidth when
// you're distributing large files. For more information about BitTorrent, see
// Amazon S3 Torrent (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html).
//
// You can get torrent only for objects that are less than 5 GB in size and
// that are not encrypted using server-side encryption with customer-provided
// encryption key.
//
// To use GET, you must have READ access to the object.
//
// The following operation is related to GetObjectTorrent:
//
//    * GetObject
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

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.GetObjectTorrentMarshaler{Input: input}.GetNamedBuildHandler())

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
