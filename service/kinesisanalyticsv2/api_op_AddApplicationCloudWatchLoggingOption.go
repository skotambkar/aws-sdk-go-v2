// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
)

const opAddApplicationCloudWatchLoggingOption = "AddApplicationCloudWatchLoggingOption"

// AddApplicationCloudWatchLoggingOptionRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds an Amazon CloudWatch log stream to monitor application configuration
// errors.
//
//    // Example sending a request using AddApplicationCloudWatchLoggingOptionRequest.
//    req := client.AddApplicationCloudWatchLoggingOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationCloudWatchLoggingOption
func (c *Client) AddApplicationCloudWatchLoggingOptionRequest(input *types.AddApplicationCloudWatchLoggingOptionInput) AddApplicationCloudWatchLoggingOptionRequest {
	op := &aws.Operation{
		Name:       opAddApplicationCloudWatchLoggingOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddApplicationCloudWatchLoggingOptionInput{}
	}

	req := c.newRequest(op, input, &types.AddApplicationCloudWatchLoggingOptionOutput{})
	return AddApplicationCloudWatchLoggingOptionRequest{Request: req, Input: input, Copy: c.AddApplicationCloudWatchLoggingOptionRequest}
}

// AddApplicationCloudWatchLoggingOptionRequest is the request type for the
// AddApplicationCloudWatchLoggingOption API operation.
type AddApplicationCloudWatchLoggingOptionRequest struct {
	*aws.Request
	Input *types.AddApplicationCloudWatchLoggingOptionInput
	Copy  func(*types.AddApplicationCloudWatchLoggingOptionInput) AddApplicationCloudWatchLoggingOptionRequest
}

// Send marshals and sends the AddApplicationCloudWatchLoggingOption API request.
func (r AddApplicationCloudWatchLoggingOptionRequest) Send(ctx context.Context) (*AddApplicationCloudWatchLoggingOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationCloudWatchLoggingOptionResponse{
		AddApplicationCloudWatchLoggingOptionOutput: r.Request.Data.(*types.AddApplicationCloudWatchLoggingOptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationCloudWatchLoggingOptionResponse is the response type for the
// AddApplicationCloudWatchLoggingOption API operation.
type AddApplicationCloudWatchLoggingOptionResponse struct {
	*types.AddApplicationCloudWatchLoggingOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationCloudWatchLoggingOption request.
func (r *AddApplicationCloudWatchLoggingOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
