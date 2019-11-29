// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

const opRemovePermission = "RemovePermission"

// RemovePermissionRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Revokes the permission of another AWS account to be able to put events to
// the specified event bus. Specify the account to revoke by the StatementId
// value that you associated with the account when you granted it permission
// with PutPermission. You can find the StatementId by using DescribeEventBus.
//
//    // Example sending a request using RemovePermissionRequest.
//    req := client.RemovePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/RemovePermission
func (c *Client) RemovePermissionRequest(input *types.RemovePermissionInput) RemovePermissionRequest {
	op := &aws.Operation{
		Name:       opRemovePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemovePermissionInput{}
	}

	req := c.newRequest(op, input, &types.RemovePermissionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemovePermissionRequest{Request: req, Input: input, Copy: c.RemovePermissionRequest}
}

// RemovePermissionRequest is the request type for the
// RemovePermission API operation.
type RemovePermissionRequest struct {
	*aws.Request
	Input *types.RemovePermissionInput
	Copy  func(*types.RemovePermissionInput) RemovePermissionRequest
}

// Send marshals and sends the RemovePermission API request.
func (r RemovePermissionRequest) Send(ctx context.Context) (*RemovePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemovePermissionResponse{
		RemovePermissionOutput: r.Request.Data.(*types.RemovePermissionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemovePermissionResponse is the response type for the
// RemovePermission API operation.
type RemovePermissionResponse struct {
	*types.RemovePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemovePermission request.
func (r *RemovePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
