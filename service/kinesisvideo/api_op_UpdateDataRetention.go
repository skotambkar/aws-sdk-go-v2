// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
)

const opUpdateDataRetention = "UpdateDataRetention"

// UpdateDataRetentionRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Increases or decreases the stream's data retention period by the value that
// you specify. To indicate whether you want to increase or decrease the data
// retention period, specify the Operation parameter in the request body. In
// the request, you must specify either the StreamName or the StreamARN.
//
// The retention period that you specify replaces the current value.
//
// This operation requires permission for the KinesisVideo:UpdateDataRetention
// action.
//
// Changing the data retention period affects the data in the stream as follows:
//
//    * If the data retention period is increased, existing data is retained
//    for the new retention period. For example, if the data retention period
//    is increased from one hour to seven hours, all existing data is retained
//    for seven hours.
//
//    * If the data retention period is decreased, existing data is retained
//    for the new retention period. For example, if the data retention period
//    is decreased from seven hours to one hour, all existing data is retained
//    for one hour, and any data older than one hour is deleted immediately.
//
//    // Example sending a request using UpdateDataRetentionRequest.
//    req := client.UpdateDataRetentionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/UpdateDataRetention
func (c *Client) UpdateDataRetentionRequest(input *types.UpdateDataRetentionInput) UpdateDataRetentionRequest {
	op := &aws.Operation{
		Name:       opUpdateDataRetention,
		HTTPMethod: "POST",
		HTTPPath:   "/updateDataRetention",
	}

	if input == nil {
		input = &types.UpdateDataRetentionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDataRetentionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateDataRetentionMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDataRetentionRequest{Request: req, Input: input, Copy: c.UpdateDataRetentionRequest}
}

// UpdateDataRetentionRequest is the request type for the
// UpdateDataRetention API operation.
type UpdateDataRetentionRequest struct {
	*aws.Request
	Input *types.UpdateDataRetentionInput
	Copy  func(*types.UpdateDataRetentionInput) UpdateDataRetentionRequest
}

// Send marshals and sends the UpdateDataRetention API request.
func (r UpdateDataRetentionRequest) Send(ctx context.Context) (*UpdateDataRetentionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataRetentionResponse{
		UpdateDataRetentionOutput: r.Request.Data.(*types.UpdateDataRetentionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataRetentionResponse is the response type for the
// UpdateDataRetention API operation.
type UpdateDataRetentionResponse struct {
	*types.UpdateDataRetentionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataRetention request.
func (r *UpdateDataRetentionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
