// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opUpdateEndpointsBatch = "UpdateEndpointsBatch"

// UpdateEndpointsBatchRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a new batch of endpoints for an application or updates the settings
// and attributes of a batch of existing endpoints for an application. You can
// also use this operation to define custom attributes (Attributes, Metrics,
// and UserAttributes properties) for a batch of endpoints.
//
//    // Example sending a request using UpdateEndpointsBatchRequest.
//    req := client.UpdateEndpointsBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateEndpointsBatch
func (c *Client) UpdateEndpointsBatchRequest(input *types.UpdateEndpointsBatchInput) UpdateEndpointsBatchRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpointsBatch,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/endpoints",
	}

	if input == nil {
		input = &types.UpdateEndpointsBatchInput{}
	}

	req := c.newRequest(op, input, &types.UpdateEndpointsBatchOutput{})
	return UpdateEndpointsBatchRequest{Request: req, Input: input, Copy: c.UpdateEndpointsBatchRequest}
}

// UpdateEndpointsBatchRequest is the request type for the
// UpdateEndpointsBatch API operation.
type UpdateEndpointsBatchRequest struct {
	*aws.Request
	Input *types.UpdateEndpointsBatchInput
	Copy  func(*types.UpdateEndpointsBatchInput) UpdateEndpointsBatchRequest
}

// Send marshals and sends the UpdateEndpointsBatch API request.
func (r UpdateEndpointsBatchRequest) Send(ctx context.Context) (*UpdateEndpointsBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointsBatchResponse{
		UpdateEndpointsBatchOutput: r.Request.Data.(*types.UpdateEndpointsBatchOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointsBatchResponse is the response type for the
// UpdateEndpointsBatch API operation.
type UpdateEndpointsBatchResponse struct {
	*types.UpdateEndpointsBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpointsBatch request.
func (r *UpdateEndpointsBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
