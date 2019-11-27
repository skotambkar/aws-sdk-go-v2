// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opStartConfigurationRecorder = "StartConfigurationRecorder"

// StartConfigurationRecorderRequest returns a request value for making API operation for
// AWS Config.
//
// Starts recording configurations of the AWS resources you have selected to
// record in your AWS account.
//
// You must have created at least one delivery channel to successfully start
// the configuration recorder.
//
//    // Example sending a request using StartConfigurationRecorderRequest.
//    req := client.StartConfigurationRecorderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/StartConfigurationRecorder
func (c *Client) StartConfigurationRecorderRequest(input *types.StartConfigurationRecorderInput) StartConfigurationRecorderRequest {
	op := &aws.Operation{
		Name:       opStartConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartConfigurationRecorderInput{}
	}

	req := c.newRequest(op, input, &types.StartConfigurationRecorderOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartConfigurationRecorderRequest{Request: req, Input: input, Copy: c.StartConfigurationRecorderRequest}
}

// StartConfigurationRecorderRequest is the request type for the
// StartConfigurationRecorder API operation.
type StartConfigurationRecorderRequest struct {
	*aws.Request
	Input *types.StartConfigurationRecorderInput
	Copy  func(*types.StartConfigurationRecorderInput) StartConfigurationRecorderRequest
}

// Send marshals and sends the StartConfigurationRecorder API request.
func (r StartConfigurationRecorderRequest) Send(ctx context.Context) (*StartConfigurationRecorderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartConfigurationRecorderResponse{
		StartConfigurationRecorderOutput: r.Request.Data.(*types.StartConfigurationRecorderOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartConfigurationRecorderResponse is the response type for the
// StartConfigurationRecorder API operation.
type StartConfigurationRecorderResponse struct {
	*types.StartConfigurationRecorderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartConfigurationRecorder request.
func (r *StartConfigurationRecorderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
