// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opStartRelationalDatabase = "StartRelationalDatabase"

// StartRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Starts a specific database from a stopped state in Amazon Lightsail. To restart
// a database, use the reboot relational database operation.
//
// The start relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using StartRelationalDatabaseRequest.
//    req := client.StartRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/StartRelationalDatabase
func (c *Client) StartRelationalDatabaseRequest(input *types.StartRelationalDatabaseInput) StartRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opStartRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &types.StartRelationalDatabaseOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartRelationalDatabaseMarshaler{Input: input}.GetNamedBuildHandler())

	return StartRelationalDatabaseRequest{Request: req, Input: input, Copy: c.StartRelationalDatabaseRequest}
}

// StartRelationalDatabaseRequest is the request type for the
// StartRelationalDatabase API operation.
type StartRelationalDatabaseRequest struct {
	*aws.Request
	Input *types.StartRelationalDatabaseInput
	Copy  func(*types.StartRelationalDatabaseInput) StartRelationalDatabaseRequest
}

// Send marshals and sends the StartRelationalDatabase API request.
func (r StartRelationalDatabaseRequest) Send(ctx context.Context) (*StartRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartRelationalDatabaseResponse{
		StartRelationalDatabaseOutput: r.Request.Data.(*types.StartRelationalDatabaseOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartRelationalDatabaseResponse is the response type for the
// StartRelationalDatabase API operation.
type StartRelationalDatabaseResponse struct {
	*types.StartRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartRelationalDatabase request.
func (r *StartRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
