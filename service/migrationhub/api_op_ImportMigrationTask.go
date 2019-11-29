// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opImportMigrationTask = "ImportMigrationTask"

// ImportMigrationTaskRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Registers a new migration task which represents a server, database, etc.,
// being migrated to AWS by a migration tool.
//
// This API is a prerequisite to calling the NotifyMigrationTaskState API as
// the migration tool must first register the migration task with Migration
// Hub.
//
//    // Example sending a request using ImportMigrationTaskRequest.
//    req := client.ImportMigrationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ImportMigrationTask
func (c *Client) ImportMigrationTaskRequest(input *types.ImportMigrationTaskInput) ImportMigrationTaskRequest {
	op := &aws.Operation{
		Name:       opImportMigrationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ImportMigrationTaskInput{}
	}

	req := c.newRequest(op, input, &types.ImportMigrationTaskOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ImportMigrationTaskMarshaler{Input: input}.GetNamedBuildHandler())

	return ImportMigrationTaskRequest{Request: req, Input: input, Copy: c.ImportMigrationTaskRequest}
}

// ImportMigrationTaskRequest is the request type for the
// ImportMigrationTask API operation.
type ImportMigrationTaskRequest struct {
	*aws.Request
	Input *types.ImportMigrationTaskInput
	Copy  func(*types.ImportMigrationTaskInput) ImportMigrationTaskRequest
}

// Send marshals and sends the ImportMigrationTask API request.
func (r ImportMigrationTaskRequest) Send(ctx context.Context) (*ImportMigrationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportMigrationTaskResponse{
		ImportMigrationTaskOutput: r.Request.Data.(*types.ImportMigrationTaskOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportMigrationTaskResponse is the response type for the
// ImportMigrationTask API operation.
type ImportMigrationTaskResponse struct {
	*types.ImportMigrationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportMigrationTask request.
func (r *ImportMigrationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
