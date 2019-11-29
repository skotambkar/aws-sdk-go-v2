// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
)

const opDescribeSharedDirectories = "DescribeSharedDirectories"

// DescribeSharedDirectoriesRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Returns the shared directories in your account.
//
//    // Example sending a request using DescribeSharedDirectoriesRequest.
//    req := client.DescribeSharedDirectoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DescribeSharedDirectories
func (c *Client) DescribeSharedDirectoriesRequest(input *types.DescribeSharedDirectoriesInput) DescribeSharedDirectoriesRequest {
	op := &aws.Operation{
		Name:       opDescribeSharedDirectories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSharedDirectoriesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSharedDirectoriesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeSharedDirectoriesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeSharedDirectoriesRequest{Request: req, Input: input, Copy: c.DescribeSharedDirectoriesRequest}
}

// DescribeSharedDirectoriesRequest is the request type for the
// DescribeSharedDirectories API operation.
type DescribeSharedDirectoriesRequest struct {
	*aws.Request
	Input *types.DescribeSharedDirectoriesInput
	Copy  func(*types.DescribeSharedDirectoriesInput) DescribeSharedDirectoriesRequest
}

// Send marshals and sends the DescribeSharedDirectories API request.
func (r DescribeSharedDirectoriesRequest) Send(ctx context.Context) (*DescribeSharedDirectoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSharedDirectoriesResponse{
		DescribeSharedDirectoriesOutput: r.Request.Data.(*types.DescribeSharedDirectoriesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSharedDirectoriesResponse is the response type for the
// DescribeSharedDirectories API operation.
type DescribeSharedDirectoriesResponse struct {
	*types.DescribeSharedDirectoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSharedDirectories request.
func (r *DescribeSharedDirectoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
