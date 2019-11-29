// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteFlowLogs = "DeleteFlowLogs"

// DeleteFlowLogsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes one or more flow logs.
//
//    // Example sending a request using DeleteFlowLogsRequest.
//    req := client.DeleteFlowLogsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteFlowLogs
func (c *Client) DeleteFlowLogsRequest(input *types.DeleteFlowLogsInput) DeleteFlowLogsRequest {
	op := &aws.Operation{
		Name:       opDeleteFlowLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteFlowLogsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteFlowLogsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteFlowLogsMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteFlowLogsRequest{Request: req, Input: input, Copy: c.DeleteFlowLogsRequest}
}

// DeleteFlowLogsRequest is the request type for the
// DeleteFlowLogs API operation.
type DeleteFlowLogsRequest struct {
	*aws.Request
	Input *types.DeleteFlowLogsInput
	Copy  func(*types.DeleteFlowLogsInput) DeleteFlowLogsRequest
}

// Send marshals and sends the DeleteFlowLogs API request.
func (r DeleteFlowLogsRequest) Send(ctx context.Context) (*DeleteFlowLogsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFlowLogsResponse{
		DeleteFlowLogsOutput: r.Request.Data.(*types.DeleteFlowLogsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFlowLogsResponse is the response type for the
// DeleteFlowLogs API operation.
type DeleteFlowLogsResponse struct {
	*types.DeleteFlowLogsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFlowLogs request.
func (r *DeleteFlowLogsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
