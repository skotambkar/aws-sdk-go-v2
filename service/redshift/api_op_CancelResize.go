// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opCancelResize = "CancelResize"

// CancelResizeRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Cancels a resize operation.
//
//    // Example sending a request using CancelResizeRequest.
//    req := client.CancelResizeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CancelResize
func (c *Client) CancelResizeRequest(input *types.CancelResizeInput) CancelResizeRequest {
	op := &aws.Operation{
		Name:       opCancelResize,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CancelResizeInput{}
	}

	req := c.newRequest(op, input, &types.CancelResizeOutput{})
	return CancelResizeRequest{Request: req, Input: input, Copy: c.CancelResizeRequest}
}

// CancelResizeRequest is the request type for the
// CancelResize API operation.
type CancelResizeRequest struct {
	*aws.Request
	Input *types.CancelResizeInput
	Copy  func(*types.CancelResizeInput) CancelResizeRequest
}

// Send marshals and sends the CancelResize API request.
func (r CancelResizeRequest) Send(ctx context.Context) (*CancelResizeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelResizeResponse{
		CancelResizeOutput: r.Request.Data.(*types.CancelResizeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelResizeResponse is the response type for the
// CancelResize API operation.
type CancelResizeResponse struct {
	*types.CancelResizeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelResize request.
func (r *CancelResizeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
