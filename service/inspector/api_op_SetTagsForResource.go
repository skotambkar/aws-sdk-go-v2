// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

const opSetTagsForResource = "SetTagsForResource"

// SetTagsForResourceRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Sets tags (key and value pairs) to the assessment template that is specified
// by the ARN of the assessment template.
//
//    // Example sending a request using SetTagsForResourceRequest.
//    req := client.SetTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/SetTagsForResource
func (c *Client) SetTagsForResourceRequest(input *types.SetTagsForResourceInput) SetTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opSetTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &types.SetTagsForResourceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetTagsForResourceRequest{Request: req, Input: input, Copy: c.SetTagsForResourceRequest}
}

// SetTagsForResourceRequest is the request type for the
// SetTagsForResource API operation.
type SetTagsForResourceRequest struct {
	*aws.Request
	Input *types.SetTagsForResourceInput
	Copy  func(*types.SetTagsForResourceInput) SetTagsForResourceRequest
}

// Send marshals and sends the SetTagsForResource API request.
func (r SetTagsForResourceRequest) Send(ctx context.Context) (*SetTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetTagsForResourceResponse{
		SetTagsForResourceOutput: r.Request.Data.(*types.SetTagsForResourceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetTagsForResourceResponse is the response type for the
// SetTagsForResource API operation.
type SetTagsForResourceResponse struct {
	*types.SetTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetTagsForResource request.
func (r *SetTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
