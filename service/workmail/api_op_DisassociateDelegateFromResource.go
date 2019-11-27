// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opDisassociateDelegateFromResource = "DisassociateDelegateFromResource"

// DisassociateDelegateFromResourceRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Removes a member from the resource's set of delegates.
//
//    // Example sending a request using DisassociateDelegateFromResourceRequest.
//    req := client.DisassociateDelegateFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DisassociateDelegateFromResource
func (c *Client) DisassociateDelegateFromResourceRequest(input *types.DisassociateDelegateFromResourceInput) DisassociateDelegateFromResourceRequest {
	op := &aws.Operation{
		Name:       opDisassociateDelegateFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateDelegateFromResourceInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateDelegateFromResourceOutput{})
	return DisassociateDelegateFromResourceRequest{Request: req, Input: input, Copy: c.DisassociateDelegateFromResourceRequest}
}

// DisassociateDelegateFromResourceRequest is the request type for the
// DisassociateDelegateFromResource API operation.
type DisassociateDelegateFromResourceRequest struct {
	*aws.Request
	Input *types.DisassociateDelegateFromResourceInput
	Copy  func(*types.DisassociateDelegateFromResourceInput) DisassociateDelegateFromResourceRequest
}

// Send marshals and sends the DisassociateDelegateFromResource API request.
func (r DisassociateDelegateFromResourceRequest) Send(ctx context.Context) (*DisassociateDelegateFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDelegateFromResourceResponse{
		DisassociateDelegateFromResourceOutput: r.Request.Data.(*types.DisassociateDelegateFromResourceOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDelegateFromResourceResponse is the response type for the
// DisassociateDelegateFromResource API operation.
type DisassociateDelegateFromResourceResponse struct {
	*types.DisassociateDelegateFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDelegateFromResource request.
func (r *DisassociateDelegateFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
