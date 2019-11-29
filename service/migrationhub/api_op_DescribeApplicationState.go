// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opDescribeApplicationState = "DescribeApplicationState"

// DescribeApplicationStateRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Gets the migration status of an application.
//
//    // Example sending a request using DescribeApplicationStateRequest.
//    req := client.DescribeApplicationStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DescribeApplicationState
func (c *Client) DescribeApplicationStateRequest(input *types.DescribeApplicationStateInput) DescribeApplicationStateRequest {
	op := &aws.Operation{
		Name:       opDescribeApplicationState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeApplicationStateInput{}
	}

	req := c.newRequest(op, input, &types.DescribeApplicationStateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeApplicationStateMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeApplicationStateRequest{Request: req, Input: input, Copy: c.DescribeApplicationStateRequest}
}

// DescribeApplicationStateRequest is the request type for the
// DescribeApplicationState API operation.
type DescribeApplicationStateRequest struct {
	*aws.Request
	Input *types.DescribeApplicationStateInput
	Copy  func(*types.DescribeApplicationStateInput) DescribeApplicationStateRequest
}

// Send marshals and sends the DescribeApplicationState API request.
func (r DescribeApplicationStateRequest) Send(ctx context.Context) (*DescribeApplicationStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationStateResponse{
		DescribeApplicationStateOutput: r.Request.Data.(*types.DescribeApplicationStateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationStateResponse is the response type for the
// DescribeApplicationState API operation.
type DescribeApplicationStateResponse struct {
	*types.DescribeApplicationStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplicationState request.
func (r *DescribeApplicationStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
