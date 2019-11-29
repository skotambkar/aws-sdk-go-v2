// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListViolationEvents = "ListViolationEvents"

// ListViolationEventsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the Device Defender security profile violations discovered during the
// given time period. You can use filters to limit the results to those alerts
// issued for a particular security profile, behavior, or thing (device).
//
//    // Example sending a request using ListViolationEventsRequest.
//    req := client.ListViolationEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListViolationEventsRequest(input *types.ListViolationEventsInput) ListViolationEventsRequest {
	op := &aws.Operation{
		Name:       opListViolationEvents,
		HTTPMethod: "GET",
		HTTPPath:   "/violation-events",
	}

	if input == nil {
		input = &types.ListViolationEventsInput{}
	}

	req := c.newRequest(op, input, &types.ListViolationEventsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListViolationEventsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListViolationEventsRequest{Request: req, Input: input, Copy: c.ListViolationEventsRequest}
}

// ListViolationEventsRequest is the request type for the
// ListViolationEvents API operation.
type ListViolationEventsRequest struct {
	*aws.Request
	Input *types.ListViolationEventsInput
	Copy  func(*types.ListViolationEventsInput) ListViolationEventsRequest
}

// Send marshals and sends the ListViolationEvents API request.
func (r ListViolationEventsRequest) Send(ctx context.Context) (*ListViolationEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListViolationEventsResponse{
		ListViolationEventsOutput: r.Request.Data.(*types.ListViolationEventsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListViolationEventsResponse is the response type for the
// ListViolationEvents API operation.
type ListViolationEventsResponse struct {
	*types.ListViolationEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListViolationEvents request.
func (r *ListViolationEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
