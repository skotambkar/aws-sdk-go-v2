// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
)

const opDeleteApplication = "DeleteApplication"

// DeleteApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes the specified application. Amazon Kinesis Analytics halts application
// execution and deletes the application, including any application artifacts
// (such as in-application streams, reference table, and application code).
//
// This operation requires permissions to perform the kinesisanalytics:DeleteApplication
// action.
//
//    // Example sending a request using DeleteApplicationRequest.
//    req := client.DeleteApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplication
func (c *Client) DeleteApplicationRequest(input *types.DeleteApplicationInput) DeleteApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteApplicationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteApplicationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteApplicationMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteApplicationRequest{Request: req, Input: input, Copy: c.DeleteApplicationRequest}
}

// DeleteApplicationRequest is the request type for the
// DeleteApplication API operation.
type DeleteApplicationRequest struct {
	*aws.Request
	Input *types.DeleteApplicationInput
	Copy  func(*types.DeleteApplicationInput) DeleteApplicationRequest
}

// Send marshals and sends the DeleteApplication API request.
func (r DeleteApplicationRequest) Send(ctx context.Context) (*DeleteApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationResponse{
		DeleteApplicationOutput: r.Request.Data.(*types.DeleteApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationResponse is the response type for the
// DeleteApplication API operation.
type DeleteApplicationResponse struct {
	*types.DeleteApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplication request.
func (r *DeleteApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
