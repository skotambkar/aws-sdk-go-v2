// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opGetBackupVaultAccessPolicy = "GetBackupVaultAccessPolicy"

// GetBackupVaultAccessPolicyRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the access policy document that is associated with the named backup
// vault.
//
//    // Example sending a request using GetBackupVaultAccessPolicyRequest.
//    req := client.GetBackupVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupVaultAccessPolicy
func (c *Client) GetBackupVaultAccessPolicyRequest(input *types.GetBackupVaultAccessPolicyInput) GetBackupVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opGetBackupVaultAccessPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-vaults/{backupVaultName}/access-policy",
	}

	if input == nil {
		input = &types.GetBackupVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &types.GetBackupVaultAccessPolicyOutput{})
	return GetBackupVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.GetBackupVaultAccessPolicyRequest}
}

// GetBackupVaultAccessPolicyRequest is the request type for the
// GetBackupVaultAccessPolicy API operation.
type GetBackupVaultAccessPolicyRequest struct {
	*aws.Request
	Input *types.GetBackupVaultAccessPolicyInput
	Copy  func(*types.GetBackupVaultAccessPolicyInput) GetBackupVaultAccessPolicyRequest
}

// Send marshals and sends the GetBackupVaultAccessPolicy API request.
func (r GetBackupVaultAccessPolicyRequest) Send(ctx context.Context) (*GetBackupVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupVaultAccessPolicyResponse{
		GetBackupVaultAccessPolicyOutput: r.Request.Data.(*types.GetBackupVaultAccessPolicyOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupVaultAccessPolicyResponse is the response type for the
// GetBackupVaultAccessPolicy API operation.
type GetBackupVaultAccessPolicyResponse struct {
	*types.GetBackupVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupVaultAccessPolicy request.
func (r *GetBackupVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
