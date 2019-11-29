// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const opSuspendProcesses = "SuspendProcesses"

// SuspendProcessesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Suspends the specified automatic scaling processes, or all processes, for
// the specified Auto Scaling group.
//
// If you suspend either the Launch or Terminate process types, it can prevent
// other process types from functioning properly.
//
// To resume processes that have been suspended, use ResumeProcesses.
//
// For more information, see Suspending and Resuming Scaling Processes (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using SuspendProcessesRequest.
//    req := client.SuspendProcessesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/SuspendProcesses
func (c *Client) SuspendProcessesRequest(input *types.SuspendProcessesInput) SuspendProcessesRequest {
	op := &aws.Operation{
		Name:       opSuspendProcesses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SuspendProcessesInput{}
	}

	req := c.newRequest(op, input, &types.SuspendProcessesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.SuspendProcessesMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SuspendProcessesRequest{Request: req, Input: input, Copy: c.SuspendProcessesRequest}
}

// SuspendProcessesRequest is the request type for the
// SuspendProcesses API operation.
type SuspendProcessesRequest struct {
	*aws.Request
	Input *types.SuspendProcessesInput
	Copy  func(*types.SuspendProcessesInput) SuspendProcessesRequest
}

// Send marshals and sends the SuspendProcesses API request.
func (r SuspendProcessesRequest) Send(ctx context.Context) (*SuspendProcessesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SuspendProcessesResponse{
		SuspendProcessesOutput: r.Request.Data.(*types.SuspendProcessesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SuspendProcessesResponse is the response type for the
// SuspendProcesses API operation.
type SuspendProcessesResponse struct {
	*types.SuspendProcessesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SuspendProcesses request.
func (r *SuspendProcessesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
