// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/backup/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opDeleteBackupVaultAccessPolicy = "DeleteBackupVaultAccessPolicy"

// DeleteBackupVaultAccessPolicyRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes the policy document that manages permissions on a backup vault.
//
//    // Example sending a request using DeleteBackupVaultAccessPolicyRequest.
//    req := client.DeleteBackupVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupVaultAccessPolicy
func (c *Client) DeleteBackupVaultAccessPolicyRequest(input *types.DeleteBackupVaultAccessPolicyInput) DeleteBackupVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupVaultAccessPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup-vaults/{backupVaultName}/access-policy",
	}

	if input == nil {
		input = &types.DeleteBackupVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBackupVaultAccessPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteBackupVaultAccessPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBackupVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.DeleteBackupVaultAccessPolicyRequest}
}

// DeleteBackupVaultAccessPolicyRequest is the request type for the
// DeleteBackupVaultAccessPolicy API operation.
type DeleteBackupVaultAccessPolicyRequest struct {
	*aws.Request
	Input *types.DeleteBackupVaultAccessPolicyInput
	Copy  func(*types.DeleteBackupVaultAccessPolicyInput) DeleteBackupVaultAccessPolicyRequest
}

// Send marshals and sends the DeleteBackupVaultAccessPolicy API request.
func (r DeleteBackupVaultAccessPolicyRequest) Send(ctx context.Context) (*DeleteBackupVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupVaultAccessPolicyResponse{
		DeleteBackupVaultAccessPolicyOutput: r.Request.Data.(*types.DeleteBackupVaultAccessPolicyOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupVaultAccessPolicyResponse is the response type for the
// DeleteBackupVaultAccessPolicy API operation.
type DeleteBackupVaultAccessPolicyResponse struct {
	*types.DeleteBackupVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupVaultAccessPolicy request.
func (r *DeleteBackupVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
