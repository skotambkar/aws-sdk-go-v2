// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opGetInventory = "GetInventory"

// GetInventoryRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Query inventory information.
//
//    // Example sending a request using GetInventoryRequest.
//    req := client.GetInventoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetInventory
func (c *Client) GetInventoryRequest(input *types.GetInventoryInput) GetInventoryRequest {
	op := &aws.Operation{
		Name:       opGetInventory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetInventoryInput{}
	}

	req := c.newRequest(op, input, &types.GetInventoryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetInventoryMarshaler{Input: input}.GetNamedBuildHandler())

	return GetInventoryRequest{Request: req, Input: input, Copy: c.GetInventoryRequest}
}

// GetInventoryRequest is the request type for the
// GetInventory API operation.
type GetInventoryRequest struct {
	*aws.Request
	Input *types.GetInventoryInput
	Copy  func(*types.GetInventoryInput) GetInventoryRequest
}

// Send marshals and sends the GetInventory API request.
func (r GetInventoryRequest) Send(ctx context.Context) (*GetInventoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInventoryResponse{
		GetInventoryOutput: r.Request.Data.(*types.GetInventoryOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInventoryResponse is the response type for the
// GetInventory API operation.
type GetInventoryResponse struct {
	*types.GetInventoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInventory request.
func (r *GetInventoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
