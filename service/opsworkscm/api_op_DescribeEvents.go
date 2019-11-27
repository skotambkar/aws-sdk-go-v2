// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
)

const opDescribeEvents = "DescribeEvents"

// DescribeEventsRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Describes events for a specified server. Results are ordered by time, with
// newest events first.
//
// This operation is synchronous.
//
// A ResourceNotFoundException is thrown when the server does not exist. A ValidationException
// is raised when parameters of the request are not valid.
//
//    // Example sending a request using DescribeEventsRequest.
//    req := client.DescribeEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DescribeEvents
func (c *Client) DescribeEventsRequest(input *types.DescribeEventsInput) DescribeEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEventsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEventsOutput{})
	return DescribeEventsRequest{Request: req, Input: input, Copy: c.DescribeEventsRequest}
}

// DescribeEventsRequest is the request type for the
// DescribeEvents API operation.
type DescribeEventsRequest struct {
	*aws.Request
	Input *types.DescribeEventsInput
	Copy  func(*types.DescribeEventsInput) DescribeEventsRequest
}

// Send marshals and sends the DescribeEvents API request.
func (r DescribeEventsRequest) Send(ctx context.Context) (*DescribeEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventsResponse{
		DescribeEventsOutput: r.Request.Data.(*types.DescribeEventsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventsResponse is the response type for the
// DescribeEvents API operation.
type DescribeEventsResponse struct {
	*types.DescribeEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvents request.
func (r *DescribeEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
