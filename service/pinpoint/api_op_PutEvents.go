// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opPutEvents = "PutEvents"

// PutEventsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a new event to record for endpoints, or creates or updates endpoint
// data that existing events are associated with.
//
//    // Example sending a request using PutEventsRequest.
//    req := client.PutEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/PutEvents
func (c *Client) PutEventsRequest(input *types.PutEventsInput) PutEventsRequest {
	op := &aws.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/events",
	}

	if input == nil {
		input = &types.PutEventsInput{}
	}

	req := c.newRequest(op, input, &types.PutEventsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PutEventsMarshaler{Input: input}.GetNamedBuildHandler())

	return PutEventsRequest{Request: req, Input: input, Copy: c.PutEventsRequest}
}

// PutEventsRequest is the request type for the
// PutEvents API operation.
type PutEventsRequest struct {
	*aws.Request
	Input *types.PutEventsInput
	Copy  func(*types.PutEventsInput) PutEventsRequest
}

// Send marshals and sends the PutEvents API request.
func (r PutEventsRequest) Send(ctx context.Context) (*PutEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEventsResponse{
		PutEventsOutput: r.Request.Data.(*types.PutEventsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEventsResponse is the response type for the
// PutEvents API operation.
type PutEventsResponse struct {
	*types.PutEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEvents request.
func (r *PutEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
