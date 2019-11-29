// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
)

const opDescribeParameterGroups = "DescribeParameterGroups"

// DescribeParameterGroupsRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Returns a list of parameter group descriptions. If a parameter group name
// is specified, the list will contain only the descriptions for that group.
//
//    // Example sending a request using DescribeParameterGroupsRequest.
//    req := client.DescribeParameterGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeParameterGroups
func (c *Client) DescribeParameterGroupsRequest(input *types.DescribeParameterGroupsInput) DescribeParameterGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeParameterGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeParameterGroupsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeParameterGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeParameterGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeParameterGroupsRequest{Request: req, Input: input, Copy: c.DescribeParameterGroupsRequest}
}

// DescribeParameterGroupsRequest is the request type for the
// DescribeParameterGroups API operation.
type DescribeParameterGroupsRequest struct {
	*aws.Request
	Input *types.DescribeParameterGroupsInput
	Copy  func(*types.DescribeParameterGroupsInput) DescribeParameterGroupsRequest
}

// Send marshals and sends the DescribeParameterGroups API request.
func (r DescribeParameterGroupsRequest) Send(ctx context.Context) (*DescribeParameterGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeParameterGroupsResponse{
		DescribeParameterGroupsOutput: r.Request.Data.(*types.DescribeParameterGroupsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeParameterGroupsResponse is the response type for the
// DescribeParameterGroups API operation.
type DescribeParameterGroupsResponse struct {
	*types.DescribeParameterGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeParameterGroups request.
func (r *DescribeParameterGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
