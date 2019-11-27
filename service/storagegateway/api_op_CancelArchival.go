// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opCancelArchival = "CancelArchival"

// CancelArchivalRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Cancels archiving of a virtual tape to the virtual tape shelf (VTS) after
// the archiving process is initiated. This operation is only supported in the
// tape gateway type.
//
//    // Example sending a request using CancelArchivalRequest.
//    req := client.CancelArchivalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CancelArchival
func (c *Client) CancelArchivalRequest(input *types.CancelArchivalInput) CancelArchivalRequest {
	op := &aws.Operation{
		Name:       opCancelArchival,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CancelArchivalInput{}
	}

	req := c.newRequest(op, input, &types.CancelArchivalOutput{})
	return CancelArchivalRequest{Request: req, Input: input, Copy: c.CancelArchivalRequest}
}

// CancelArchivalRequest is the request type for the
// CancelArchival API operation.
type CancelArchivalRequest struct {
	*aws.Request
	Input *types.CancelArchivalInput
	Copy  func(*types.CancelArchivalInput) CancelArchivalRequest
}

// Send marshals and sends the CancelArchival API request.
func (r CancelArchivalRequest) Send(ctx context.Context) (*CancelArchivalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelArchivalResponse{
		CancelArchivalOutput: r.Request.Data.(*types.CancelArchivalOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelArchivalResponse is the response type for the
// CancelArchival API operation.
type CancelArchivalResponse struct {
	*types.CancelArchivalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelArchival request.
func (r *CancelArchivalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
