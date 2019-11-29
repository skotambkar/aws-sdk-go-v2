// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opGetEventSourceMapping = "GetEventSourceMapping"

// GetEventSourceMappingRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns details about an event source mapping. You can get the identifier
// of a mapping from the output of ListEventSourceMappings.
//
//    // Example sending a request using GetEventSourceMappingRequest.
//    req := client.GetEventSourceMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetEventSourceMapping
func (c *Client) GetEventSourceMappingRequest(input *types.GetEventSourceMappingInput) GetEventSourceMappingRequest {
	op := &aws.Operation{
		Name:       opGetEventSourceMapping,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/event-source-mappings/{UUID}",
	}

	if input == nil {
		input = &types.GetEventSourceMappingInput{}
	}

	req := c.newRequest(op, input, &types.GetEventSourceMappingOutput{})
	return GetEventSourceMappingRequest{Request: req, Input: input, Copy: c.GetEventSourceMappingRequest}
}

// GetEventSourceMappingRequest is the request type for the
// GetEventSourceMapping API operation.
type GetEventSourceMappingRequest struct {
	*aws.Request
	Input *types.GetEventSourceMappingInput
	Copy  func(*types.GetEventSourceMappingInput) GetEventSourceMappingRequest
}

// Send marshals and sends the GetEventSourceMapping API request.
func (r GetEventSourceMappingRequest) Send(ctx context.Context) (*GetEventSourceMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEventSourceMappingResponse{
		GetEventSourceMappingOutput: r.Request.Data.(*types.GetEventSourceMappingOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEventSourceMappingResponse is the response type for the
// GetEventSourceMapping API operation.
type GetEventSourceMappingResponse struct {
	*types.GetEventSourceMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEventSourceMapping request.
func (r *GetEventSourceMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
