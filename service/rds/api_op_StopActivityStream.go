// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opStopActivityStream = "StopActivityStream"

// StopActivityStreamRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Stops a database activity stream that was started using the AWS console,
// the start-activity-stream AWS CLI command, or the StartActivityStream action.
//
// For more information, see Database Activity Streams (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html)
// in the Amazon Aurora User Guide.
//
//    // Example sending a request using StopActivityStreamRequest.
//    req := client.StopActivityStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StopActivityStream
func (c *Client) StopActivityStreamRequest(input *types.StopActivityStreamInput) StopActivityStreamRequest {
	op := &aws.Operation{
		Name:       opStopActivityStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopActivityStreamInput{}
	}

	req := c.newRequest(op, input, &types.StopActivityStreamOutput{})
	return StopActivityStreamRequest{Request: req, Input: input, Copy: c.StopActivityStreamRequest}
}

// StopActivityStreamRequest is the request type for the
// StopActivityStream API operation.
type StopActivityStreamRequest struct {
	*aws.Request
	Input *types.StopActivityStreamInput
	Copy  func(*types.StopActivityStreamInput) StopActivityStreamRequest
}

// Send marshals and sends the StopActivityStream API request.
func (r StopActivityStreamRequest) Send(ctx context.Context) (*StopActivityStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopActivityStreamResponse{
		StopActivityStreamOutput: r.Request.Data.(*types.StopActivityStreamOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopActivityStreamResponse is the response type for the
// StopActivityStream API operation.
type StopActivityStreamResponse struct {
	*types.StopActivityStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopActivityStream request.
func (r *StopActivityStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
