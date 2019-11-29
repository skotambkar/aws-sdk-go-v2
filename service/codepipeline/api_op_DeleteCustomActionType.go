// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

const opDeleteCustomActionType = "DeleteCustomActionType"

// DeleteCustomActionTypeRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Marks a custom action as deleted. PollForJobs for the custom action fails
// after the action is marked for deletion. Used for custom actions only.
//
// To re-create a custom action after it has been deleted you must use a string
// in the version field that has never been used before. This string can be
// an incremented version number, for example. To restore a deleted custom action,
// use a JSON file that is identical to the deleted action, including the original
// string in the version field.
//
//    // Example sending a request using DeleteCustomActionTypeRequest.
//    req := client.DeleteCustomActionTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeleteCustomActionType
func (c *Client) DeleteCustomActionTypeRequest(input *types.DeleteCustomActionTypeInput) DeleteCustomActionTypeRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomActionType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCustomActionTypeInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCustomActionTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCustomActionTypeRequest{Request: req, Input: input, Copy: c.DeleteCustomActionTypeRequest}
}

// DeleteCustomActionTypeRequest is the request type for the
// DeleteCustomActionType API operation.
type DeleteCustomActionTypeRequest struct {
	*aws.Request
	Input *types.DeleteCustomActionTypeInput
	Copy  func(*types.DeleteCustomActionTypeInput) DeleteCustomActionTypeRequest
}

// Send marshals and sends the DeleteCustomActionType API request.
func (r DeleteCustomActionTypeRequest) Send(ctx context.Context) (*DeleteCustomActionTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomActionTypeResponse{
		DeleteCustomActionTypeOutput: r.Request.Data.(*types.DeleteCustomActionTypeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomActionTypeResponse is the response type for the
// DeleteCustomActionType API operation.
type DeleteCustomActionTypeResponse struct {
	*types.DeleteCustomActionTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomActionType request.
func (r *DeleteCustomActionTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
