// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDeleteChapCredentials = "DeleteChapCredentials"

// DeleteChapCredentialsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes Challenge-Handshake Authentication Protocol (CHAP) credentials for
// a specified iSCSI target and initiator pair. This operation is supported
// in volume and tape gateway types.
//
//    // Example sending a request using DeleteChapCredentialsRequest.
//    req := client.DeleteChapCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteChapCredentials
func (c *Client) DeleteChapCredentialsRequest(input *types.DeleteChapCredentialsInput) DeleteChapCredentialsRequest {
	op := &aws.Operation{
		Name:       opDeleteChapCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteChapCredentialsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteChapCredentialsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteChapCredentialsMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteChapCredentialsRequest{Request: req, Input: input, Copy: c.DeleteChapCredentialsRequest}
}

// DeleteChapCredentialsRequest is the request type for the
// DeleteChapCredentials API operation.
type DeleteChapCredentialsRequest struct {
	*aws.Request
	Input *types.DeleteChapCredentialsInput
	Copy  func(*types.DeleteChapCredentialsInput) DeleteChapCredentialsRequest
}

// Send marshals and sends the DeleteChapCredentials API request.
func (r DeleteChapCredentialsRequest) Send(ctx context.Context) (*DeleteChapCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteChapCredentialsResponse{
		DeleteChapCredentialsOutput: r.Request.Data.(*types.DeleteChapCredentialsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteChapCredentialsResponse is the response type for the
// DeleteChapCredentials API operation.
type DeleteChapCredentialsResponse struct {
	*types.DeleteChapCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteChapCredentials request.
func (r *DeleteChapCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
