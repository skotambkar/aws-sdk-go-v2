// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opRecordHandlerProgress = "RecordHandlerProgress"

// RecordHandlerProgressRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Reports progress of a resource handler to CloudFormation.
//
// Reserved for use by the CloudFormation CLI (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
// Do not use this API in your code.
//
//    // Example sending a request using RecordHandlerProgressRequest.
//    req := client.RecordHandlerProgressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RecordHandlerProgress
func (c *Client) RecordHandlerProgressRequest(input *types.RecordHandlerProgressInput) RecordHandlerProgressRequest {
	op := &aws.Operation{
		Name:       opRecordHandlerProgress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RecordHandlerProgressInput{}
	}

	req := c.newRequest(op, input, &types.RecordHandlerProgressOutput{})
	return RecordHandlerProgressRequest{Request: req, Input: input, Copy: c.RecordHandlerProgressRequest}
}

// RecordHandlerProgressRequest is the request type for the
// RecordHandlerProgress API operation.
type RecordHandlerProgressRequest struct {
	*aws.Request
	Input *types.RecordHandlerProgressInput
	Copy  func(*types.RecordHandlerProgressInput) RecordHandlerProgressRequest
}

// Send marshals and sends the RecordHandlerProgress API request.
func (r RecordHandlerProgressRequest) Send(ctx context.Context) (*RecordHandlerProgressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RecordHandlerProgressResponse{
		RecordHandlerProgressOutput: r.Request.Data.(*types.RecordHandlerProgressOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RecordHandlerProgressResponse is the response type for the
// RecordHandlerProgress API operation.
type RecordHandlerProgressResponse struct {
	*types.RecordHandlerProgressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RecordHandlerProgress request.
func (r *RecordHandlerProgressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
