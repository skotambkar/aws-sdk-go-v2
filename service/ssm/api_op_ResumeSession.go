// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opResumeSession = "ResumeSession"

// ResumeSessionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Reconnects a session to an instance after it has been disconnected. Connections
// can be resumed for disconnected sessions, but not terminated sessions.
//
// This command is primarily for use by client machines to automatically reconnect
// during intermittent network issues. It is not intended for any other use.
//
//    // Example sending a request using ResumeSessionRequest.
//    req := client.ResumeSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ResumeSession
func (c *Client) ResumeSessionRequest(input *types.ResumeSessionInput) ResumeSessionRequest {
	op := &aws.Operation{
		Name:       opResumeSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResumeSessionInput{}
	}

	req := c.newRequest(op, input, &types.ResumeSessionOutput{})
	return ResumeSessionRequest{Request: req, Input: input, Copy: c.ResumeSessionRequest}
}

// ResumeSessionRequest is the request type for the
// ResumeSession API operation.
type ResumeSessionRequest struct {
	*aws.Request
	Input *types.ResumeSessionInput
	Copy  func(*types.ResumeSessionInput) ResumeSessionRequest
}

// Send marshals and sends the ResumeSession API request.
func (r ResumeSessionRequest) Send(ctx context.Context) (*ResumeSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResumeSessionResponse{
		ResumeSessionOutput: r.Request.Data.(*types.ResumeSessionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResumeSessionResponse is the response type for the
// ResumeSession API operation.
type ResumeSessionResponse struct {
	*types.ResumeSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResumeSession request.
func (r *ResumeSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
