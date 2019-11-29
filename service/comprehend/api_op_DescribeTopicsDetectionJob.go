// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opDescribeTopicsDetectionJob = "DescribeTopicsDetectionJob"

// DescribeTopicsDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets the properties associated with a topic detection job. Use this operation
// to get the status of a detection job.
//
//    // Example sending a request using DescribeTopicsDetectionJobRequest.
//    req := client.DescribeTopicsDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeTopicsDetectionJob
func (c *Client) DescribeTopicsDetectionJobRequest(input *types.DescribeTopicsDetectionJobInput) DescribeTopicsDetectionJobRequest {
	op := &aws.Operation{
		Name:       opDescribeTopicsDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTopicsDetectionJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTopicsDetectionJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeTopicsDetectionJobMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeTopicsDetectionJobRequest{Request: req, Input: input, Copy: c.DescribeTopicsDetectionJobRequest}
}

// DescribeTopicsDetectionJobRequest is the request type for the
// DescribeTopicsDetectionJob API operation.
type DescribeTopicsDetectionJobRequest struct {
	*aws.Request
	Input *types.DescribeTopicsDetectionJobInput
	Copy  func(*types.DescribeTopicsDetectionJobInput) DescribeTopicsDetectionJobRequest
}

// Send marshals and sends the DescribeTopicsDetectionJob API request.
func (r DescribeTopicsDetectionJobRequest) Send(ctx context.Context) (*DescribeTopicsDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTopicsDetectionJobResponse{
		DescribeTopicsDetectionJobOutput: r.Request.Data.(*types.DescribeTopicsDetectionJobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTopicsDetectionJobResponse is the response type for the
// DescribeTopicsDetectionJob API operation.
type DescribeTopicsDetectionJobResponse struct {
	*types.DescribeTopicsDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTopicsDetectionJob request.
func (r *DescribeTopicsDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
