// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
)

const opResolveCase = "ResolveCase"

// ResolveCaseRequest returns a request value for making API operation for
// AWS Support.
//
// Takes a caseId and returns the initial state of the case along with the state
// of the case after the call to ResolveCase completed.
//
//    // Example sending a request using ResolveCaseRequest.
//    req := client.ResolveCaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/ResolveCase
func (c *Client) ResolveCaseRequest(input *types.ResolveCaseInput) ResolveCaseRequest {
	op := &aws.Operation{
		Name:       opResolveCase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResolveCaseInput{}
	}

	req := c.newRequest(op, input, &types.ResolveCaseOutput{})
	return ResolveCaseRequest{Request: req, Input: input, Copy: c.ResolveCaseRequest}
}

// ResolveCaseRequest is the request type for the
// ResolveCase API operation.
type ResolveCaseRequest struct {
	*aws.Request
	Input *types.ResolveCaseInput
	Copy  func(*types.ResolveCaseInput) ResolveCaseRequest
}

// Send marshals and sends the ResolveCase API request.
func (r ResolveCaseRequest) Send(ctx context.Context) (*ResolveCaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResolveCaseResponse{
		ResolveCaseOutput: r.Request.Data.(*types.ResolveCaseOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResolveCaseResponse is the response type for the
// ResolveCase API operation.
type ResolveCaseResponse struct {
	*types.ResolveCaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResolveCase request.
func (r *ResolveCaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
