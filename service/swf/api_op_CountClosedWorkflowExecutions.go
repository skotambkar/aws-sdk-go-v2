// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opCountClosedWorkflowExecutions = "CountClosedWorkflowExecutions"

// CountClosedWorkflowExecutionsRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the number of closed workflow executions within the given domain
// that meet the specified filtering criteria.
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
//    the appropriate keys. tagFilter.tag: String constraint. The key is swf:tagFilter.tag.
//    typeFilter.name: String constraint. The key is swf:typeFilter.name. typeFilter.version:
//    String constraint. The key is swf:typeFilter.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using CountClosedWorkflowExecutionsRequest.
//    req := client.CountClosedWorkflowExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CountClosedWorkflowExecutionsRequest(input *types.CountClosedWorkflowExecutionsInput) CountClosedWorkflowExecutionsRequest {
	op := &aws.Operation{
		Name:       opCountClosedWorkflowExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CountClosedWorkflowExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.CountClosedWorkflowExecutionsOutput{})
	return CountClosedWorkflowExecutionsRequest{Request: req, Input: input, Copy: c.CountClosedWorkflowExecutionsRequest}
}

// CountClosedWorkflowExecutionsRequest is the request type for the
// CountClosedWorkflowExecutions API operation.
type CountClosedWorkflowExecutionsRequest struct {
	*aws.Request
	Input *types.CountClosedWorkflowExecutionsInput
	Copy  func(*types.CountClosedWorkflowExecutionsInput) CountClosedWorkflowExecutionsRequest
}

// Send marshals and sends the CountClosedWorkflowExecutions API request.
func (r CountClosedWorkflowExecutionsRequest) Send(ctx context.Context) (*CountClosedWorkflowExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CountClosedWorkflowExecutionsResponse{
		CountClosedWorkflowExecutionsOutput: r.Request.Data.(*types.CountClosedWorkflowExecutionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CountClosedWorkflowExecutionsResponse is the response type for the
// CountClosedWorkflowExecutions API operation.
type CountClosedWorkflowExecutionsResponse struct {
	*types.CountClosedWorkflowExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CountClosedWorkflowExecutions request.
func (r *CountClosedWorkflowExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
