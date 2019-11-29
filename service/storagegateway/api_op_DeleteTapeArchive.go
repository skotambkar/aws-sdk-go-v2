// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDeleteTapeArchive = "DeleteTapeArchive"

// DeleteTapeArchiveRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the specified virtual tape from the virtual tape shelf (VTS). This
// operation is only supported in the tape gateway type.
//
//    // Example sending a request using DeleteTapeArchiveRequest.
//    req := client.DeleteTapeArchiveRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteTapeArchive
func (c *Client) DeleteTapeArchiveRequest(input *types.DeleteTapeArchiveInput) DeleteTapeArchiveRequest {
	op := &aws.Operation{
		Name:       opDeleteTapeArchive,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTapeArchiveInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTapeArchiveOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteTapeArchiveMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteTapeArchiveRequest{Request: req, Input: input, Copy: c.DeleteTapeArchiveRequest}
}

// DeleteTapeArchiveRequest is the request type for the
// DeleteTapeArchive API operation.
type DeleteTapeArchiveRequest struct {
	*aws.Request
	Input *types.DeleteTapeArchiveInput
	Copy  func(*types.DeleteTapeArchiveInput) DeleteTapeArchiveRequest
}

// Send marshals and sends the DeleteTapeArchive API request.
func (r DeleteTapeArchiveRequest) Send(ctx context.Context) (*DeleteTapeArchiveResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTapeArchiveResponse{
		DeleteTapeArchiveOutput: r.Request.Data.(*types.DeleteTapeArchiveOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTapeArchiveResponse is the response type for the
// DeleteTapeArchive API operation.
type DeleteTapeArchiveResponse struct {
	*types.DeleteTapeArchiveOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTapeArchive request.
func (r *DeleteTapeArchiveResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
