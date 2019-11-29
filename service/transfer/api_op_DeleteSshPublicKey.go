// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
)

const opDeleteSshPublicKey = "DeleteSshPublicKey"

// DeleteSshPublicKeyRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Deletes a user's Secure Shell (SSH) public key.
//
// No response is returned from this operation.
//
//    // Example sending a request using DeleteSshPublicKeyRequest.
//    req := client.DeleteSshPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DeleteSshPublicKey
func (c *Client) DeleteSshPublicKeyRequest(input *types.DeleteSshPublicKeyInput) DeleteSshPublicKeyRequest {
	op := &aws.Operation{
		Name:       opDeleteSshPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSshPublicKeyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSshPublicKeyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSshPublicKeyRequest{Request: req, Input: input, Copy: c.DeleteSshPublicKeyRequest}
}

// DeleteSshPublicKeyRequest is the request type for the
// DeleteSshPublicKey API operation.
type DeleteSshPublicKeyRequest struct {
	*aws.Request
	Input *types.DeleteSshPublicKeyInput
	Copy  func(*types.DeleteSshPublicKeyInput) DeleteSshPublicKeyRequest
}

// Send marshals and sends the DeleteSshPublicKey API request.
func (r DeleteSshPublicKeyRequest) Send(ctx context.Context) (*DeleteSshPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSshPublicKeyResponse{
		DeleteSshPublicKeyOutput: r.Request.Data.(*types.DeleteSshPublicKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSshPublicKeyResponse is the response type for the
// DeleteSshPublicKey API operation.
type DeleteSshPublicKeyResponse struct {
	*types.DeleteSshPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSshPublicKey request.
func (r *DeleteSshPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
