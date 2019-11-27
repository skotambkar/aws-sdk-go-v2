// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opGetInventorySchema = "GetInventorySchema"

// GetInventorySchemaRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Return a list of inventory type names for the account, or return a list of
// attribute names for a specific Inventory item type.
//
//    // Example sending a request using GetInventorySchemaRequest.
//    req := client.GetInventorySchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetInventorySchema
func (c *Client) GetInventorySchemaRequest(input *types.GetInventorySchemaInput) GetInventorySchemaRequest {
	op := &aws.Operation{
		Name:       opGetInventorySchema,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetInventorySchemaInput{}
	}

	req := c.newRequest(op, input, &types.GetInventorySchemaOutput{})
	return GetInventorySchemaRequest{Request: req, Input: input, Copy: c.GetInventorySchemaRequest}
}

// GetInventorySchemaRequest is the request type for the
// GetInventorySchema API operation.
type GetInventorySchemaRequest struct {
	*aws.Request
	Input *types.GetInventorySchemaInput
	Copy  func(*types.GetInventorySchemaInput) GetInventorySchemaRequest
}

// Send marshals and sends the GetInventorySchema API request.
func (r GetInventorySchemaRequest) Send(ctx context.Context) (*GetInventorySchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInventorySchemaResponse{
		GetInventorySchemaOutput: r.Request.Data.(*types.GetInventorySchemaOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInventorySchemaResponse is the response type for the
// GetInventorySchema API operation.
type GetInventorySchemaResponse struct {
	*types.GetInventorySchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInventorySchema request.
func (r *GetInventorySchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
