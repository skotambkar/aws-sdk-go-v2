// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDescribeChannel = "DescribeChannel"

// DescribeChannelRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a channel.
//
//    // Example sending a request using DescribeChannelRequest.
//    req := client.DescribeChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeChannel
func (c *Client) DescribeChannelRequest(input *types.DescribeChannelInput) DescribeChannelRequest {
	op := &aws.Operation{
		Name:       opDescribeChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/channels/{channelName}",
	}

	if input == nil {
		input = &types.DescribeChannelInput{}
	}

	req := c.newRequest(op, input, &types.DescribeChannelOutput{})
	return DescribeChannelRequest{Request: req, Input: input, Copy: c.DescribeChannelRequest}
}

// DescribeChannelRequest is the request type for the
// DescribeChannel API operation.
type DescribeChannelRequest struct {
	*aws.Request
	Input *types.DescribeChannelInput
	Copy  func(*types.DescribeChannelInput) DescribeChannelRequest
}

// Send marshals and sends the DescribeChannel API request.
func (r DescribeChannelRequest) Send(ctx context.Context) (*DescribeChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChannelResponse{
		DescribeChannelOutput: r.Request.Data.(*types.DescribeChannelOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChannelResponse is the response type for the
// DescribeChannel API operation.
type DescribeChannelResponse struct {
	*types.DescribeChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChannel request.
func (r *DescribeChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
