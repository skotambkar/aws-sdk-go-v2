// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opDeleteBGPPeer = "DeleteBGPPeer"

// DeleteBGPPeerRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes the specified BGP peer on the specified virtual interface with the
// specified customer address and ASN.
//
// You cannot delete the last BGP peer from a virtual interface.
//
//    // Example sending a request using DeleteBGPPeerRequest.
//    req := client.DeleteBGPPeerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteBGPPeer
func (c *Client) DeleteBGPPeerRequest(input *types.DeleteBGPPeerInput) DeleteBGPPeerRequest {
	op := &aws.Operation{
		Name:       opDeleteBGPPeer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteBGPPeerInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBGPPeerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteBGPPeerMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteBGPPeerRequest{Request: req, Input: input, Copy: c.DeleteBGPPeerRequest}
}

// DeleteBGPPeerRequest is the request type for the
// DeleteBGPPeer API operation.
type DeleteBGPPeerRequest struct {
	*aws.Request
	Input *types.DeleteBGPPeerInput
	Copy  func(*types.DeleteBGPPeerInput) DeleteBGPPeerRequest
}

// Send marshals and sends the DeleteBGPPeer API request.
func (r DeleteBGPPeerRequest) Send(ctx context.Context) (*DeleteBGPPeerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBGPPeerResponse{
		DeleteBGPPeerOutput: r.Request.Data.(*types.DeleteBGPPeerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBGPPeerResponse is the response type for the
// DeleteBGPPeer API operation.
type DeleteBGPPeerResponse struct {
	*types.DeleteBGPPeerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBGPPeer request.
func (r *DeleteBGPPeerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
