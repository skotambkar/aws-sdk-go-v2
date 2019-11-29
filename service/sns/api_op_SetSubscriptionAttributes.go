// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opSetSubscriptionAttributes = "SetSubscriptionAttributes"

// SetSubscriptionAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Allows a subscription owner to set an attribute of the subscription to a
// new value.
//
//    // Example sending a request using SetSubscriptionAttributesRequest.
//    req := client.SetSubscriptionAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetSubscriptionAttributes
func (c *Client) SetSubscriptionAttributesRequest(input *types.SetSubscriptionAttributesInput) SetSubscriptionAttributesRequest {
	op := &aws.Operation{
		Name:       opSetSubscriptionAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetSubscriptionAttributesInput{}
	}

	req := c.newRequest(op, input, &types.SetSubscriptionAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetSubscriptionAttributesRequest{Request: req, Input: input, Copy: c.SetSubscriptionAttributesRequest}
}

// SetSubscriptionAttributesRequest is the request type for the
// SetSubscriptionAttributes API operation.
type SetSubscriptionAttributesRequest struct {
	*aws.Request
	Input *types.SetSubscriptionAttributesInput
	Copy  func(*types.SetSubscriptionAttributesInput) SetSubscriptionAttributesRequest
}

// Send marshals and sends the SetSubscriptionAttributes API request.
func (r SetSubscriptionAttributesRequest) Send(ctx context.Context) (*SetSubscriptionAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSubscriptionAttributesResponse{
		SetSubscriptionAttributesOutput: r.Request.Data.(*types.SetSubscriptionAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSubscriptionAttributesResponse is the response type for the
// SetSubscriptionAttributes API operation.
type SetSubscriptionAttributesResponse struct {
	*types.SetSubscriptionAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSubscriptionAttributes request.
func (r *SetSubscriptionAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
