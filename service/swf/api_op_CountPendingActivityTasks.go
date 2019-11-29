// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opCountPendingActivityTasks = "CountPendingActivityTasks"

// CountPendingActivityTasksRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the estimated number of activity tasks in the specified task list.
// The count returned is an approximation and isn't guaranteed to be exact.
// If you specify a task list that no activity task was ever scheduled in then
// 0 is returned.
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
//    * Constrain the taskList.name parameter by using a Condition element with
//    the swf:taskList.name key to allow the action to access only certain task
//    lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using CountPendingActivityTasksRequest.
//    req := client.CountPendingActivityTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CountPendingActivityTasksRequest(input *types.CountPendingActivityTasksInput) CountPendingActivityTasksRequest {
	op := &aws.Operation{
		Name:       opCountPendingActivityTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CountPendingActivityTasksInput{}
	}

	req := c.newRequest(op, input, &types.CountPendingActivityTasksOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CountPendingActivityTasksMarshaler{Input: input}.GetNamedBuildHandler())

	return CountPendingActivityTasksRequest{Request: req, Input: input, Copy: c.CountPendingActivityTasksRequest}
}

// CountPendingActivityTasksRequest is the request type for the
// CountPendingActivityTasks API operation.
type CountPendingActivityTasksRequest struct {
	*aws.Request
	Input *types.CountPendingActivityTasksInput
	Copy  func(*types.CountPendingActivityTasksInput) CountPendingActivityTasksRequest
}

// Send marshals and sends the CountPendingActivityTasks API request.
func (r CountPendingActivityTasksRequest) Send(ctx context.Context) (*CountPendingActivityTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CountPendingActivityTasksResponse{
		CountPendingActivityTasksOutput: r.Request.Data.(*types.CountPendingActivityTasksOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CountPendingActivityTasksResponse is the response type for the
// CountPendingActivityTasks API operation.
type CountPendingActivityTasksResponse struct {
	*types.CountPendingActivityTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CountPendingActivityTasks request.
func (r *CountPendingActivityTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
