// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDeleteMitigationAction = "DeleteMitigationAction"

// DeleteMitigationActionRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a defined mitigation action from your AWS account.
//
//    // Example sending a request using DeleteMitigationActionRequest.
//    req := client.DeleteMitigationActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteMitigationActionRequest(input *types.DeleteMitigationActionInput) DeleteMitigationActionRequest {
	op := &aws.Operation{
		Name:       opDeleteMitigationAction,
		HTTPMethod: "DELETE",
		HTTPPath:   "/mitigationactions/actions/{actionName}",
	}

	if input == nil {
		input = &types.DeleteMitigationActionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteMitigationActionOutput{})
	return DeleteMitigationActionRequest{Request: req, Input: input, Copy: c.DeleteMitigationActionRequest}
}

// DeleteMitigationActionRequest is the request type for the
// DeleteMitigationAction API operation.
type DeleteMitigationActionRequest struct {
	*aws.Request
	Input *types.DeleteMitigationActionInput
	Copy  func(*types.DeleteMitigationActionInput) DeleteMitigationActionRequest
}

// Send marshals and sends the DeleteMitigationAction API request.
func (r DeleteMitigationActionRequest) Send(ctx context.Context) (*DeleteMitigationActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMitigationActionResponse{
		DeleteMitigationActionOutput: r.Request.Data.(*types.DeleteMitigationActionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMitigationActionResponse is the response type for the
// DeleteMitigationAction API operation.
type DeleteMitigationActionResponse struct {
	*types.DeleteMitigationActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMitigationAction request.
func (r *DeleteMitigationActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
