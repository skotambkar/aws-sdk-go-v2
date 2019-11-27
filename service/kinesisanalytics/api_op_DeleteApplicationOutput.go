// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
)

const opDeleteApplicationOutput = "DeleteApplicationOutput"

// DeleteApplicationOutputRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes output destination configuration from your application configuration.
// Amazon Kinesis Analytics will no longer write data from the corresponding
// in-application stream to the external output destination.
//
// This operation requires permissions to perform the kinesisanalytics:DeleteApplicationOutput
// action.
//
//    // Example sending a request using DeleteApplicationOutputRequest.
//    req := client.DeleteApplicationOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationOutput
func (c *Client) DeleteApplicationOutputRequest(input *types.DeleteApplicationOutputInput) DeleteApplicationOutputRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationOutput,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteApplicationOutputInput{}
	}

	req := c.newRequest(op, input, &types.DeleteApplicationOutputOutput{})
	return DeleteApplicationOutputRequest{Request: req, Input: input, Copy: c.DeleteApplicationOutputRequest}
}

// DeleteApplicationOutputRequest is the request type for the
// DeleteApplicationOutput API operation.
type DeleteApplicationOutputRequest struct {
	*aws.Request
	Input *types.DeleteApplicationOutputInput
	Copy  func(*types.DeleteApplicationOutputInput) DeleteApplicationOutputRequest
}

// Send marshals and sends the DeleteApplicationOutput API request.
func (r DeleteApplicationOutputRequest) Send(ctx context.Context) (*DeleteApplicationOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationOutputResponse{
		DeleteApplicationOutputOutput: r.Request.Data.(*types.DeleteApplicationOutputOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationOutputResponse is the response type for the
// DeleteApplicationOutput API operation.
type DeleteApplicationOutputResponse struct {
	*types.DeleteApplicationOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationOutput request.
func (r *DeleteApplicationOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
