// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opSetTopicAttributes = "SetTopicAttributes"

// SetTopicAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Allows a topic owner to set an attribute of the topic to a new value.
//
//    // Example sending a request using SetTopicAttributesRequest.
//    req := client.SetTopicAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetTopicAttributes
func (c *Client) SetTopicAttributesRequest(input *types.SetTopicAttributesInput) SetTopicAttributesRequest {
	op := &aws.Operation{
		Name:       opSetTopicAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetTopicAttributesInput{}
	}

	req := c.newRequest(op, input, &types.SetTopicAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetTopicAttributesRequest{Request: req, Input: input, Copy: c.SetTopicAttributesRequest}
}

// SetTopicAttributesRequest is the request type for the
// SetTopicAttributes API operation.
type SetTopicAttributesRequest struct {
	*aws.Request
	Input *types.SetTopicAttributesInput
	Copy  func(*types.SetTopicAttributesInput) SetTopicAttributesRequest
}

// Send marshals and sends the SetTopicAttributes API request.
func (r SetTopicAttributesRequest) Send(ctx context.Context) (*SetTopicAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetTopicAttributesResponse{
		SetTopicAttributesOutput: r.Request.Data.(*types.SetTopicAttributesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetTopicAttributesResponse is the response type for the
// SetTopicAttributes API operation.
type SetTopicAttributesResponse struct {
	*types.SetTopicAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetTopicAttributes request.
func (r *SetTopicAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
