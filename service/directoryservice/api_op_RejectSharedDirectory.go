// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
)

const opRejectSharedDirectory = "RejectSharedDirectory"

// RejectSharedDirectoryRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Rejects a directory sharing request that was sent from the directory owner
// account.
//
//    // Example sending a request using RejectSharedDirectoryRequest.
//    req := client.RejectSharedDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/RejectSharedDirectory
func (c *Client) RejectSharedDirectoryRequest(input *types.RejectSharedDirectoryInput) RejectSharedDirectoryRequest {
	op := &aws.Operation{
		Name:       opRejectSharedDirectory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RejectSharedDirectoryInput{}
	}

	req := c.newRequest(op, input, &types.RejectSharedDirectoryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RejectSharedDirectoryMarshaler{Input: input}.GetNamedBuildHandler())

	return RejectSharedDirectoryRequest{Request: req, Input: input, Copy: c.RejectSharedDirectoryRequest}
}

// RejectSharedDirectoryRequest is the request type for the
// RejectSharedDirectory API operation.
type RejectSharedDirectoryRequest struct {
	*aws.Request
	Input *types.RejectSharedDirectoryInput
	Copy  func(*types.RejectSharedDirectoryInput) RejectSharedDirectoryRequest
}

// Send marshals and sends the RejectSharedDirectory API request.
func (r RejectSharedDirectoryRequest) Send(ctx context.Context) (*RejectSharedDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectSharedDirectoryResponse{
		RejectSharedDirectoryOutput: r.Request.Data.(*types.RejectSharedDirectoryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectSharedDirectoryResponse is the response type for the
// RejectSharedDirectory API operation.
type RejectSharedDirectoryResponse struct {
	*types.RejectSharedDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectSharedDirectory request.
func (r *RejectSharedDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
