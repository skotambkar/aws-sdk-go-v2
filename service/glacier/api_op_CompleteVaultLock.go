// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

const opCompleteVaultLock = "CompleteVaultLock"

// CompleteVaultLockRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation completes the vault locking process by transitioning the vault
// lock from the InProgress state to the Locked state, which causes the vault
// lock policy to become unchangeable. A vault lock is put into the InProgress
// state by calling InitiateVaultLock. You can obtain the state of the vault
// lock by calling GetVaultLock. For more information about the vault locking
// process, Amazon Glacier Vault Lock (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html).
//
// This operation is idempotent. This request is always successful if the vault
// lock is in the Locked state and the provided lock ID matches the lock ID
// originally used to lock the vault.
//
// If an invalid lock ID is passed in the request when the vault lock is in
// the Locked state, the operation returns an AccessDeniedException error. If
// an invalid lock ID is passed in the request when the vault lock is in the
// InProgress state, the operation throws an InvalidParameter error.
//
//    // Example sending a request using CompleteVaultLockRequest.
//    req := client.CompleteVaultLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CompleteVaultLockRequest(input *types.CompleteVaultLockInput) CompleteVaultLockRequest {
	op := &aws.Operation{
		Name:       opCompleteVaultLock,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/lock-policy/{lockId}",
	}

	if input == nil {
		input = &types.CompleteVaultLockInput{}
	}

	req := c.newRequest(op, input, &types.CompleteVaultLockOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CompleteVaultLockMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CompleteVaultLockRequest{Request: req, Input: input, Copy: c.CompleteVaultLockRequest}
}

// CompleteVaultLockRequest is the request type for the
// CompleteVaultLock API operation.
type CompleteVaultLockRequest struct {
	*aws.Request
	Input *types.CompleteVaultLockInput
	Copy  func(*types.CompleteVaultLockInput) CompleteVaultLockRequest
}

// Send marshals and sends the CompleteVaultLock API request.
func (r CompleteVaultLockRequest) Send(ctx context.Context) (*CompleteVaultLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompleteVaultLockResponse{
		CompleteVaultLockOutput: r.Request.Data.(*types.CompleteVaultLockOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompleteVaultLockResponse is the response type for the
// CompleteVaultLock API operation.
type CompleteVaultLockResponse struct {
	*types.CompleteVaultLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompleteVaultLock request.
func (r *CompleteVaultLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
