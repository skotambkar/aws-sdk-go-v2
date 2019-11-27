// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opUndeprecateWorkflowType = "UndeprecateWorkflowType"

// UndeprecateWorkflowTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Undeprecates a previously deprecated workflow type. After a workflow type
// has been undeprecated, you can create new executions of that type.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. workflowType.name: String constraint. The key is
//    swf:workflowType.name. workflowType.version: String constraint. The key
//    is swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using UndeprecateWorkflowTypeRequest.
//    req := client.UndeprecateWorkflowTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UndeprecateWorkflowTypeRequest(input *types.UndeprecateWorkflowTypeInput) UndeprecateWorkflowTypeRequest {
	op := &aws.Operation{
		Name:       opUndeprecateWorkflowType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UndeprecateWorkflowTypeInput{}
	}

	req := c.newRequest(op, input, &types.UndeprecateWorkflowTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UndeprecateWorkflowTypeRequest{Request: req, Input: input, Copy: c.UndeprecateWorkflowTypeRequest}
}

// UndeprecateWorkflowTypeRequest is the request type for the
// UndeprecateWorkflowType API operation.
type UndeprecateWorkflowTypeRequest struct {
	*aws.Request
	Input *types.UndeprecateWorkflowTypeInput
	Copy  func(*types.UndeprecateWorkflowTypeInput) UndeprecateWorkflowTypeRequest
}

// Send marshals and sends the UndeprecateWorkflowType API request.
func (r UndeprecateWorkflowTypeRequest) Send(ctx context.Context) (*UndeprecateWorkflowTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UndeprecateWorkflowTypeResponse{
		UndeprecateWorkflowTypeOutput: r.Request.Data.(*types.UndeprecateWorkflowTypeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UndeprecateWorkflowTypeResponse is the response type for the
// UndeprecateWorkflowType API operation.
type UndeprecateWorkflowTypeResponse struct {
	*types.UndeprecateWorkflowTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UndeprecateWorkflowType request.
func (r *UndeprecateWorkflowTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
