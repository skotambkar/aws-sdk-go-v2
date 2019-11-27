// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opListFunctionDefinitions = "ListFunctionDefinitions"

// ListFunctionDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of Lambda function definitions.
//
//    // Example sending a request using ListFunctionDefinitionsRequest.
//    req := client.ListFunctionDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListFunctionDefinitions
func (c *Client) ListFunctionDefinitionsRequest(input *types.ListFunctionDefinitionsInput) ListFunctionDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListFunctionDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/functions",
	}

	if input == nil {
		input = &types.ListFunctionDefinitionsInput{}
	}

	req := c.newRequest(op, input, &types.ListFunctionDefinitionsOutput{})
	return ListFunctionDefinitionsRequest{Request: req, Input: input, Copy: c.ListFunctionDefinitionsRequest}
}

// ListFunctionDefinitionsRequest is the request type for the
// ListFunctionDefinitions API operation.
type ListFunctionDefinitionsRequest struct {
	*aws.Request
	Input *types.ListFunctionDefinitionsInput
	Copy  func(*types.ListFunctionDefinitionsInput) ListFunctionDefinitionsRequest
}

// Send marshals and sends the ListFunctionDefinitions API request.
func (r ListFunctionDefinitionsRequest) Send(ctx context.Context) (*ListFunctionDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFunctionDefinitionsResponse{
		ListFunctionDefinitionsOutput: r.Request.Data.(*types.ListFunctionDefinitionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFunctionDefinitionsResponse is the response type for the
// ListFunctionDefinitions API operation.
type ListFunctionDefinitionsResponse struct {
	*types.ListFunctionDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFunctionDefinitions request.
func (r *ListFunctionDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
