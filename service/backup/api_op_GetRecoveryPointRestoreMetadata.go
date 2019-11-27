// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opGetRecoveryPointRestoreMetadata = "GetRecoveryPointRestoreMetadata"

// GetRecoveryPointRestoreMetadataRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns two sets of metadata key-value pairs. The first set lists the metadata
// that the recovery point was created with. The second set lists the metadata
// key-value pairs that are required to restore the recovery point.
//
// These sets can be the same, or the restore metadata set can contain different
// values if the target service to be restored has changed since the recovery
// point was created and now requires additional or different information in
// order to be restored.
//
//    // Example sending a request using GetRecoveryPointRestoreMetadataRequest.
//    req := client.GetRecoveryPointRestoreMetadataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetRecoveryPointRestoreMetadata
func (c *Client) GetRecoveryPointRestoreMetadataRequest(input *types.GetRecoveryPointRestoreMetadataInput) GetRecoveryPointRestoreMetadataRequest {
	op := &aws.Operation{
		Name:       opGetRecoveryPointRestoreMetadata,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-vaults/{backupVaultName}/recovery-points/{recoveryPointArn}/restore-metadata",
	}

	if input == nil {
		input = &types.GetRecoveryPointRestoreMetadataInput{}
	}

	req := c.newRequest(op, input, &types.GetRecoveryPointRestoreMetadataOutput{})
	return GetRecoveryPointRestoreMetadataRequest{Request: req, Input: input, Copy: c.GetRecoveryPointRestoreMetadataRequest}
}

// GetRecoveryPointRestoreMetadataRequest is the request type for the
// GetRecoveryPointRestoreMetadata API operation.
type GetRecoveryPointRestoreMetadataRequest struct {
	*aws.Request
	Input *types.GetRecoveryPointRestoreMetadataInput
	Copy  func(*types.GetRecoveryPointRestoreMetadataInput) GetRecoveryPointRestoreMetadataRequest
}

// Send marshals and sends the GetRecoveryPointRestoreMetadata API request.
func (r GetRecoveryPointRestoreMetadataRequest) Send(ctx context.Context) (*GetRecoveryPointRestoreMetadataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecoveryPointRestoreMetadataResponse{
		GetRecoveryPointRestoreMetadataOutput: r.Request.Data.(*types.GetRecoveryPointRestoreMetadataOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecoveryPointRestoreMetadataResponse is the response type for the
// GetRecoveryPointRestoreMetadata API operation.
type GetRecoveryPointRestoreMetadataResponse struct {
	*types.GetRecoveryPointRestoreMetadataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecoveryPointRestoreMetadata request.
func (r *GetRecoveryPointRestoreMetadataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
