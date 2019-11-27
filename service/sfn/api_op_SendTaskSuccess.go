// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
)

const opSendTaskSuccess = "SendTaskSuccess"

// SendTaskSuccessRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Used by activity workers and task states using the callback (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern to report that the task identified by the taskToken completed successfully.
//
//    // Example sending a request using SendTaskSuccessRequest.
//    req := client.SendTaskSuccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/SendTaskSuccess
func (c *Client) SendTaskSuccessRequest(input *types.SendTaskSuccessInput) SendTaskSuccessRequest {
	op := &aws.Operation{
		Name:       opSendTaskSuccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SendTaskSuccessInput{}
	}

	req := c.newRequest(op, input, &types.SendTaskSuccessOutput{})
	return SendTaskSuccessRequest{Request: req, Input: input, Copy: c.SendTaskSuccessRequest}
}

// SendTaskSuccessRequest is the request type for the
// SendTaskSuccess API operation.
type SendTaskSuccessRequest struct {
	*aws.Request
	Input *types.SendTaskSuccessInput
	Copy  func(*types.SendTaskSuccessInput) SendTaskSuccessRequest
}

// Send marshals and sends the SendTaskSuccess API request.
func (r SendTaskSuccessRequest) Send(ctx context.Context) (*SendTaskSuccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendTaskSuccessResponse{
		SendTaskSuccessOutput: r.Request.Data.(*types.SendTaskSuccessOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendTaskSuccessResponse is the response type for the
// SendTaskSuccess API operation.
type SendTaskSuccessResponse struct {
	*types.SendTaskSuccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendTaskSuccess request.
func (r *SendTaskSuccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
