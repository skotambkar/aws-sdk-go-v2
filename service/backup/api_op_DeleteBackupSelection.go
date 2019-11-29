// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opDeleteBackupSelection = "DeleteBackupSelection"

// DeleteBackupSelectionRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes the resource selection associated with a backup plan that is specified
// by the SelectionId.
//
//    // Example sending a request using DeleteBackupSelectionRequest.
//    req := client.DeleteBackupSelectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupSelection
func (c *Client) DeleteBackupSelectionRequest(input *types.DeleteBackupSelectionInput) DeleteBackupSelectionRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupSelection,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/{selectionId}",
	}

	if input == nil {
		input = &types.DeleteBackupSelectionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBackupSelectionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBackupSelectionRequest{Request: req, Input: input, Copy: c.DeleteBackupSelectionRequest}
}

// DeleteBackupSelectionRequest is the request type for the
// DeleteBackupSelection API operation.
type DeleteBackupSelectionRequest struct {
	*aws.Request
	Input *types.DeleteBackupSelectionInput
	Copy  func(*types.DeleteBackupSelectionInput) DeleteBackupSelectionRequest
}

// Send marshals and sends the DeleteBackupSelection API request.
func (r DeleteBackupSelectionRequest) Send(ctx context.Context) (*DeleteBackupSelectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupSelectionResponse{
		DeleteBackupSelectionOutput: r.Request.Data.(*types.DeleteBackupSelectionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupSelectionResponse is the response type for the
// DeleteBackupSelection API operation.
type DeleteBackupSelectionResponse struct {
	*types.DeleteBackupSelectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupSelection request.
func (r *DeleteBackupSelectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
