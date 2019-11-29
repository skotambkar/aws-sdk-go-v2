// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
)

const opDeleteApplicationInputProcessingConfiguration = "DeleteApplicationInputProcessingConfiguration"

// DeleteApplicationInputProcessingConfigurationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes an InputProcessingConfiguration (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html)
// from an input.
//
//    // Example sending a request using DeleteApplicationInputProcessingConfigurationRequest.
//    req := client.DeleteApplicationInputProcessingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationInputProcessingConfiguration
func (c *Client) DeleteApplicationInputProcessingConfigurationRequest(input *types.DeleteApplicationInputProcessingConfigurationInput) DeleteApplicationInputProcessingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationInputProcessingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteApplicationInputProcessingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteApplicationInputProcessingConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteApplicationInputProcessingConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteApplicationInputProcessingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteApplicationInputProcessingConfigurationRequest}
}

// DeleteApplicationInputProcessingConfigurationRequest is the request type for the
// DeleteApplicationInputProcessingConfiguration API operation.
type DeleteApplicationInputProcessingConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteApplicationInputProcessingConfigurationInput
	Copy  func(*types.DeleteApplicationInputProcessingConfigurationInput) DeleteApplicationInputProcessingConfigurationRequest
}

// Send marshals and sends the DeleteApplicationInputProcessingConfiguration API request.
func (r DeleteApplicationInputProcessingConfigurationRequest) Send(ctx context.Context) (*DeleteApplicationInputProcessingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationInputProcessingConfigurationResponse{
		DeleteApplicationInputProcessingConfigurationOutput: r.Request.Data.(*types.DeleteApplicationInputProcessingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationInputProcessingConfigurationResponse is the response type for the
// DeleteApplicationInputProcessingConfiguration API operation.
type DeleteApplicationInputProcessingConfigurationResponse struct {
	*types.DeleteApplicationInputProcessingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationInputProcessingConfiguration request.
func (r *DeleteApplicationInputProcessingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
