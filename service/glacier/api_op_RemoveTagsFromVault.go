// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

const opRemoveTagsFromVault = "RemoveTagsFromVault"

// RemoveTagsFromVaultRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation removes one or more tags from the set of tags attached to
// a vault. For more information about tags, see Tagging Amazon S3 Glacier Resources
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/tagging.html). This
// operation is idempotent. The operation will be successful, even if there
// are no tags attached to the vault.
//
//    // Example sending a request using RemoveTagsFromVaultRequest.
//    req := client.RemoveTagsFromVaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RemoveTagsFromVaultRequest(input *types.RemoveTagsFromVaultInput) RemoveTagsFromVaultRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromVault,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/tags?operation=remove",
	}

	if input == nil {
		input = &types.RemoveTagsFromVaultInput{}
	}

	req := c.newRequest(op, input, &types.RemoveTagsFromVaultOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveTagsFromVaultRequest{Request: req, Input: input, Copy: c.RemoveTagsFromVaultRequest}
}

// RemoveTagsFromVaultRequest is the request type for the
// RemoveTagsFromVault API operation.
type RemoveTagsFromVaultRequest struct {
	*aws.Request
	Input *types.RemoveTagsFromVaultInput
	Copy  func(*types.RemoveTagsFromVaultInput) RemoveTagsFromVaultRequest
}

// Send marshals and sends the RemoveTagsFromVault API request.
func (r RemoveTagsFromVaultRequest) Send(ctx context.Context) (*RemoveTagsFromVaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromVaultResponse{
		RemoveTagsFromVaultOutput: r.Request.Data.(*types.RemoveTagsFromVaultOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromVaultResponse is the response type for the
// RemoveTagsFromVault API operation.
type RemoveTagsFromVaultResponse struct {
	*types.RemoveTagsFromVaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromVault request.
func (r *RemoveTagsFromVaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
