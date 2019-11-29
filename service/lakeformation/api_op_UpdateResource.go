// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opUpdateResource = "UpdateResource"

// UpdateResourceRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Updates the data access role used for vending access to the given (registered)
// resource in AWS Lake Formation.
//
//    // Example sending a request using UpdateResourceRequest.
//    req := client.UpdateResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/UpdateResource
func (c *Client) UpdateResourceRequest(input *types.UpdateResourceInput) UpdateResourceRequest {
	op := &aws.Operation{
		Name:       opUpdateResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateResourceInput{}
	}

	req := c.newRequest(op, input, &types.UpdateResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateResourceRequest{Request: req, Input: input, Copy: c.UpdateResourceRequest}
}

// UpdateResourceRequest is the request type for the
// UpdateResource API operation.
type UpdateResourceRequest struct {
	*aws.Request
	Input *types.UpdateResourceInput
	Copy  func(*types.UpdateResourceInput) UpdateResourceRequest
}

// Send marshals and sends the UpdateResource API request.
func (r UpdateResourceRequest) Send(ctx context.Context) (*UpdateResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateResourceResponse{
		UpdateResourceOutput: r.Request.Data.(*types.UpdateResourceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateResourceResponse is the response type for the
// UpdateResource API operation.
type UpdateResourceResponse struct {
	*types.UpdateResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateResource request.
func (r *UpdateResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
