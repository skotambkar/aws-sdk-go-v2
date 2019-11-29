// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opGetBackupPlan = "GetBackupPlan"

// GetBackupPlanRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the body of a backup plan in JSON format, in addition to plan metadata.
//
//    // Example sending a request using GetBackupPlanRequest.
//    req := client.GetBackupPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupPlan
func (c *Client) GetBackupPlanRequest(input *types.GetBackupPlanInput) GetBackupPlanRequest {
	op := &aws.Operation{
		Name:       opGetBackupPlan,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/",
	}

	if input == nil {
		input = &types.GetBackupPlanInput{}
	}

	req := c.newRequest(op, input, &types.GetBackupPlanOutput{})
	return GetBackupPlanRequest{Request: req, Input: input, Copy: c.GetBackupPlanRequest}
}

// GetBackupPlanRequest is the request type for the
// GetBackupPlan API operation.
type GetBackupPlanRequest struct {
	*aws.Request
	Input *types.GetBackupPlanInput
	Copy  func(*types.GetBackupPlanInput) GetBackupPlanRequest
}

// Send marshals and sends the GetBackupPlan API request.
func (r GetBackupPlanRequest) Send(ctx context.Context) (*GetBackupPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupPlanResponse{
		GetBackupPlanOutput: r.Request.Data.(*types.GetBackupPlanOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupPlanResponse is the response type for the
// GetBackupPlan API operation.
type GetBackupPlanResponse struct {
	*types.GetBackupPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupPlan request.
func (r *GetBackupPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
