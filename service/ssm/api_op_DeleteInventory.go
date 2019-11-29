// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDeleteInventory = "DeleteInventory"

// DeleteInventoryRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Delete a custom inventory type, or the data associated with a custom Inventory
// type. Deleting a custom inventory type is also referred to as deleting a
// custom inventory schema.
//
//    // Example sending a request using DeleteInventoryRequest.
//    req := client.DeleteInventoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteInventory
func (c *Client) DeleteInventoryRequest(input *types.DeleteInventoryInput) DeleteInventoryRequest {
	op := &aws.Operation{
		Name:       opDeleteInventory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteInventoryInput{}
	}

	req := c.newRequest(op, input, &types.DeleteInventoryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteInventoryMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteInventoryRequest{Request: req, Input: input, Copy: c.DeleteInventoryRequest}
}

// DeleteInventoryRequest is the request type for the
// DeleteInventory API operation.
type DeleteInventoryRequest struct {
	*aws.Request
	Input *types.DeleteInventoryInput
	Copy  func(*types.DeleteInventoryInput) DeleteInventoryRequest
}

// Send marshals and sends the DeleteInventory API request.
func (r DeleteInventoryRequest) Send(ctx context.Context) (*DeleteInventoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInventoryResponse{
		DeleteInventoryOutput: r.Request.Data.(*types.DeleteInventoryOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInventoryResponse is the response type for the
// DeleteInventory API operation.
type DeleteInventoryResponse struct {
	*types.DeleteInventoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInventory request.
func (r *DeleteInventoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
