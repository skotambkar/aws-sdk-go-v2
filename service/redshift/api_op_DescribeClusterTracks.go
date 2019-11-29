// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDescribeClusterTracks = "DescribeClusterTracks"

// DescribeClusterTracksRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of all the available maintenance tracks.
//
//    // Example sending a request using DescribeClusterTracksRequest.
//    req := client.DescribeClusterTracksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterTracks
func (c *Client) DescribeClusterTracksRequest(input *types.DescribeClusterTracksInput) DescribeClusterTracksRequest {
	op := &aws.Operation{
		Name:       opDescribeClusterTracks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeClusterTracksInput{}
	}

	req := c.newRequest(op, input, &types.DescribeClusterTracksOutput{})
	return DescribeClusterTracksRequest{Request: req, Input: input, Copy: c.DescribeClusterTracksRequest}
}

// DescribeClusterTracksRequest is the request type for the
// DescribeClusterTracks API operation.
type DescribeClusterTracksRequest struct {
	*aws.Request
	Input *types.DescribeClusterTracksInput
	Copy  func(*types.DescribeClusterTracksInput) DescribeClusterTracksRequest
}

// Send marshals and sends the DescribeClusterTracks API request.
func (r DescribeClusterTracksRequest) Send(ctx context.Context) (*DescribeClusterTracksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterTracksResponse{
		DescribeClusterTracksOutput: r.Request.Data.(*types.DescribeClusterTracksOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClusterTracksResponse is the response type for the
// DescribeClusterTracks API operation.
type DescribeClusterTracksResponse struct {
	*types.DescribeClusterTracksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusterTracks request.
func (r *DescribeClusterTracksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
