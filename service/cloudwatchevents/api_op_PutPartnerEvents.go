// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

const opPutPartnerEvents = "PutPartnerEvents"

// PutPartnerEventsRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// This is used by SaaS partners to write events to a customer's partner event
// bus.
//
// AWS customers do not use this operation. Instead, AWS customers can use PutEvents
// to write custom events from their own applications to an event bus.
//
//    // Example sending a request using PutPartnerEventsRequest.
//    req := client.PutPartnerEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/PutPartnerEvents
func (c *Client) PutPartnerEventsRequest(input *types.PutPartnerEventsInput) PutPartnerEventsRequest {
	op := &aws.Operation{
		Name:       opPutPartnerEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutPartnerEventsInput{}
	}

	req := c.newRequest(op, input, &types.PutPartnerEventsOutput{})
	return PutPartnerEventsRequest{Request: req, Input: input, Copy: c.PutPartnerEventsRequest}
}

// PutPartnerEventsRequest is the request type for the
// PutPartnerEvents API operation.
type PutPartnerEventsRequest struct {
	*aws.Request
	Input *types.PutPartnerEventsInput
	Copy  func(*types.PutPartnerEventsInput) PutPartnerEventsRequest
}

// Send marshals and sends the PutPartnerEvents API request.
func (r PutPartnerEventsRequest) Send(ctx context.Context) (*PutPartnerEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPartnerEventsResponse{
		PutPartnerEventsOutput: r.Request.Data.(*types.PutPartnerEventsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPartnerEventsResponse is the response type for the
// PutPartnerEvents API operation.
type PutPartnerEventsResponse struct {
	*types.PutPartnerEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPartnerEvents request.
func (r *PutPartnerEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
