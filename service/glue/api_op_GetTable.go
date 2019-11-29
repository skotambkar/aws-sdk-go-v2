// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opGetTable = "GetTable"

// GetTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the Table definition in a Data Catalog for a specified table.
//
//    // Example sending a request using GetTableRequest.
//    req := client.GetTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTable
func (c *Client) GetTableRequest(input *types.GetTableInput) GetTableRequest {
	op := &aws.Operation{
		Name:       opGetTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetTableInput{}
	}

	req := c.newRequest(op, input, &types.GetTableOutput{})
	return GetTableRequest{Request: req, Input: input, Copy: c.GetTableRequest}
}

// GetTableRequest is the request type for the
// GetTable API operation.
type GetTableRequest struct {
	*aws.Request
	Input *types.GetTableInput
	Copy  func(*types.GetTableInput) GetTableRequest
}

// Send marshals and sends the GetTable API request.
func (r GetTableRequest) Send(ctx context.Context) (*GetTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTableResponse{
		GetTableOutput: r.Request.Data.(*types.GetTableOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTableResponse is the response type for the
// GetTable API operation.
type GetTableResponse struct {
	*types.GetTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTable request.
func (r *GetTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
