// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const opCreateGlobalTable = "CreateGlobalTable"

// CreateGlobalTableRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Creates a global table from an existing table. A global table creates a replication
// relationship between two or more DynamoDB tables with the same table name
// in the provided Regions.
//
// If you want to add a new replica table to a global table, each of the following
// conditions must be true:
//
//    * The table must have the same primary key as all of the other replicas.
//
//    * The table must have the same name as all of the other replicas.
//
//    * The table must have DynamoDB Streams enabled, with the stream containing
//    both the new and the old images of the item.
//
//    * None of the replica tables in the global table can contain any data.
//
// If global secondary indexes are specified, then the following conditions
// must also be met:
//
//    * The global secondary indexes must have the same name.
//
//    * The global secondary indexes must have the same hash key and sort key
//    (if present).
//
// Write capacity settings should be set consistently across your replica tables
// and secondary indexes. DynamoDB strongly recommends enabling auto scaling
// to manage the write capacity settings for all of your global tables replicas
// and indexes.
//
// If you prefer to manage write capacity settings manually, you should provision
// equal replicated write capacity units to your replica tables. You should
// also provision equal replicated write capacity units to matching secondary
// indexes across your global table.
//
//    // Example sending a request using CreateGlobalTableRequest.
//    req := client.CreateGlobalTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/CreateGlobalTable
func (c *Client) CreateGlobalTableRequest(input *types.CreateGlobalTableInput) CreateGlobalTableRequest {
	op := &aws.Operation{
		Name:       opCreateGlobalTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateGlobalTableInput{}
	}

	req := c.newRequest(op, input, &types.CreateGlobalTableOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateGlobalTableMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateGlobalTableRequest{Request: req, Input: input, Copy: c.CreateGlobalTableRequest}
}

// CreateGlobalTableRequest is the request type for the
// CreateGlobalTable API operation.
type CreateGlobalTableRequest struct {
	*aws.Request
	Input *types.CreateGlobalTableInput
	Copy  func(*types.CreateGlobalTableInput) CreateGlobalTableRequest
}

// Send marshals and sends the CreateGlobalTable API request.
func (r CreateGlobalTableRequest) Send(ctx context.Context) (*CreateGlobalTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGlobalTableResponse{
		CreateGlobalTableOutput: r.Request.Data.(*types.CreateGlobalTableOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGlobalTableResponse is the response type for the
// CreateGlobalTable API operation.
type CreateGlobalTableResponse struct {
	*types.CreateGlobalTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGlobalTable request.
func (r *CreateGlobalTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
