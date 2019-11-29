// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opDescribeMigrationTask = "DescribeMigrationTask"

// DescribeMigrationTaskRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Retrieves a list of all attributes associated with a specific migration task.
//
//    // Example sending a request using DescribeMigrationTaskRequest.
//    req := client.DescribeMigrationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DescribeMigrationTask
func (c *Client) DescribeMigrationTaskRequest(input *types.DescribeMigrationTaskInput) DescribeMigrationTaskRequest {
	op := &aws.Operation{
		Name:       opDescribeMigrationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeMigrationTaskInput{}
	}

	req := c.newRequest(op, input, &types.DescribeMigrationTaskOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeMigrationTaskMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeMigrationTaskRequest{Request: req, Input: input, Copy: c.DescribeMigrationTaskRequest}
}

// DescribeMigrationTaskRequest is the request type for the
// DescribeMigrationTask API operation.
type DescribeMigrationTaskRequest struct {
	*aws.Request
	Input *types.DescribeMigrationTaskInput
	Copy  func(*types.DescribeMigrationTaskInput) DescribeMigrationTaskRequest
}

// Send marshals and sends the DescribeMigrationTask API request.
func (r DescribeMigrationTaskRequest) Send(ctx context.Context) (*DescribeMigrationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMigrationTaskResponse{
		DescribeMigrationTaskOutput: r.Request.Data.(*types.DescribeMigrationTaskOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMigrationTaskResponse is the response type for the
// DescribeMigrationTask API operation.
type DescribeMigrationTaskResponse struct {
	*types.DescribeMigrationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMigrationTask request.
func (r *DescribeMigrationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
