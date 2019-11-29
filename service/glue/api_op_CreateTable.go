// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opCreateTable = "CreateTable"

// CreateTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new table definition in the Data Catalog.
//
//    // Example sending a request using CreateTableRequest.
//    req := client.CreateTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateTable
func (c *Client) CreateTableRequest(input *types.CreateTableInput) CreateTableRequest {
	op := &aws.Operation{
		Name:       opCreateTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTableInput{}
	}

	req := c.newRequest(op, input, &types.CreateTableOutput{})
	return CreateTableRequest{Request: req, Input: input, Copy: c.CreateTableRequest}
}

// CreateTableRequest is the request type for the
// CreateTable API operation.
type CreateTableRequest struct {
	*aws.Request
	Input *types.CreateTableInput
	Copy  func(*types.CreateTableInput) CreateTableRequest
}

// Send marshals and sends the CreateTable API request.
func (r CreateTableRequest) Send(ctx context.Context) (*CreateTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTableResponse{
		CreateTableOutput: r.Request.Data.(*types.CreateTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTableResponse is the response type for the
// CreateTable API operation.
type CreateTableResponse struct {
	*types.CreateTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTable request.
func (r *CreateTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
