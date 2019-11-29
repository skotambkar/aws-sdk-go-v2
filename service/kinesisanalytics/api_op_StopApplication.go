// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
)

const opStopApplication = "StopApplication"

// StopApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Stops the application from processing input data. You can stop an application
// only if it is in the running state. You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
// operation to find the application state. After the application is stopped,
// Amazon Kinesis Analytics stops reading data from the input, the application
// stops processing data, and there is no output written to the destination.
//
// This operation requires permissions to perform the kinesisanalytics:StopApplication
// action.
//
//    // Example sending a request using StopApplicationRequest.
//    req := client.StopApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/StopApplication
func (c *Client) StopApplicationRequest(input *types.StopApplicationInput) StopApplicationRequest {
	op := &aws.Operation{
		Name:       opStopApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopApplicationInput{}
	}

	req := c.newRequest(op, input, &types.StopApplicationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StopApplicationMarshaler{Input: input}.GetNamedBuildHandler())

	return StopApplicationRequest{Request: req, Input: input, Copy: c.StopApplicationRequest}
}

// StopApplicationRequest is the request type for the
// StopApplication API operation.
type StopApplicationRequest struct {
	*aws.Request
	Input *types.StopApplicationInput
	Copy  func(*types.StopApplicationInput) StopApplicationRequest
}

// Send marshals and sends the StopApplication API request.
func (r StopApplicationRequest) Send(ctx context.Context) (*StopApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopApplicationResponse{
		StopApplicationOutput: r.Request.Data.(*types.StopApplicationOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopApplicationResponse is the response type for the
// StopApplication API operation.
type StopApplicationResponse struct {
	*types.StopApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopApplication request.
func (r *StopApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
