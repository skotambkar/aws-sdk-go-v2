// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDescribeSessions = "DescribeSessions"

// DescribeSessionsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes the streaming sessions for a specified stack
// and fleet. If a UserId is provided for the stack and fleet, only streaming
// sessions for that user are described. If an authentication type is not provided,
// the default is to authenticate users using a streaming URL.
//
//    // Example sending a request using DescribeSessionsRequest.
//    req := client.DescribeSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeSessions
func (c *Client) DescribeSessionsRequest(input *types.DescribeSessionsInput) DescribeSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribeSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSessionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSessionsOutput{})
	return DescribeSessionsRequest{Request: req, Input: input, Copy: c.DescribeSessionsRequest}
}

// DescribeSessionsRequest is the request type for the
// DescribeSessions API operation.
type DescribeSessionsRequest struct {
	*aws.Request
	Input *types.DescribeSessionsInput
	Copy  func(*types.DescribeSessionsInput) DescribeSessionsRequest
}

// Send marshals and sends the DescribeSessions API request.
func (r DescribeSessionsRequest) Send(ctx context.Context) (*DescribeSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSessionsResponse{
		DescribeSessionsOutput: r.Request.Data.(*types.DescribeSessionsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSessionsResponse is the response type for the
// DescribeSessions API operation.
type DescribeSessionsResponse struct {
	*types.DescribeSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSessions request.
func (r *DescribeSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
