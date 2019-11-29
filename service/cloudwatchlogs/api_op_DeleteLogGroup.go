// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

const opDeleteLogGroup = "DeleteLogGroup"

// DeleteLogGroupRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes the specified log group and permanently deletes all the archived
// log events associated with the log group.
//
//    // Example sending a request using DeleteLogGroupRequest.
//    req := client.DeleteLogGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteLogGroup
func (c *Client) DeleteLogGroupRequest(input *types.DeleteLogGroupInput) DeleteLogGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteLogGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteLogGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteLogGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteLogGroupRequest{Request: req, Input: input, Copy: c.DeleteLogGroupRequest}
}

// DeleteLogGroupRequest is the request type for the
// DeleteLogGroup API operation.
type DeleteLogGroupRequest struct {
	*aws.Request
	Input *types.DeleteLogGroupInput
	Copy  func(*types.DeleteLogGroupInput) DeleteLogGroupRequest
}

// Send marshals and sends the DeleteLogGroup API request.
func (r DeleteLogGroupRequest) Send(ctx context.Context) (*DeleteLogGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLogGroupResponse{
		DeleteLogGroupOutput: r.Request.Data.(*types.DeleteLogGroupOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLogGroupResponse is the response type for the
// DeleteLogGroup API operation.
type DeleteLogGroupResponse struct {
	*types.DeleteLogGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLogGroup request.
func (r *DeleteLogGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
