// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDescribeMaintenanceWindowExecutionTaskInvocations = "DescribeMaintenanceWindowExecutionTaskInvocations"

// DescribeMaintenanceWindowExecutionTaskInvocationsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the individual task executions (one per target) for a particular
// task run as part of a maintenance window execution.
//
//    // Example sending a request using DescribeMaintenanceWindowExecutionTaskInvocationsRequest.
//    req := client.DescribeMaintenanceWindowExecutionTaskInvocationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindowExecutionTaskInvocations
func (c *Client) DescribeMaintenanceWindowExecutionTaskInvocationsRequest(input *types.DescribeMaintenanceWindowExecutionTaskInvocationsInput) DescribeMaintenanceWindowExecutionTaskInvocationsRequest {
	op := &aws.Operation{
		Name:       opDescribeMaintenanceWindowExecutionTaskInvocations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeMaintenanceWindowExecutionTaskInvocationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeMaintenanceWindowExecutionTaskInvocationsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeMaintenanceWindowExecutionTaskInvocationsRequest{Request: req, Input: input, Copy: c.DescribeMaintenanceWindowExecutionTaskInvocationsRequest}
}

// DescribeMaintenanceWindowExecutionTaskInvocationsRequest is the request type for the
// DescribeMaintenanceWindowExecutionTaskInvocations API operation.
type DescribeMaintenanceWindowExecutionTaskInvocationsRequest struct {
	*aws.Request
	Input *types.DescribeMaintenanceWindowExecutionTaskInvocationsInput
	Copy  func(*types.DescribeMaintenanceWindowExecutionTaskInvocationsInput) DescribeMaintenanceWindowExecutionTaskInvocationsRequest
}

// Send marshals and sends the DescribeMaintenanceWindowExecutionTaskInvocations API request.
func (r DescribeMaintenanceWindowExecutionTaskInvocationsRequest) Send(ctx context.Context) (*DescribeMaintenanceWindowExecutionTaskInvocationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMaintenanceWindowExecutionTaskInvocationsResponse{
		DescribeMaintenanceWindowExecutionTaskInvocationsOutput: r.Request.Data.(*types.DescribeMaintenanceWindowExecutionTaskInvocationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMaintenanceWindowExecutionTaskInvocationsResponse is the response type for the
// DescribeMaintenanceWindowExecutionTaskInvocations API operation.
type DescribeMaintenanceWindowExecutionTaskInvocationsResponse struct {
	*types.DescribeMaintenanceWindowExecutionTaskInvocationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMaintenanceWindowExecutionTaskInvocations request.
func (r *DescribeMaintenanceWindowExecutionTaskInvocationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
