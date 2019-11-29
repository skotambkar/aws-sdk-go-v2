// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opDescribeEntitiesDetectionJob = "DescribeEntitiesDetectionJob"

// DescribeEntitiesDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets the properties associated with an entities detection job. Use this operation
// to get the status of a detection job.
//
//    // Example sending a request using DescribeEntitiesDetectionJobRequest.
//    req := client.DescribeEntitiesDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntitiesDetectionJob
func (c *Client) DescribeEntitiesDetectionJobRequest(input *types.DescribeEntitiesDetectionJobInput) DescribeEntitiesDetectionJobRequest {
	op := &aws.Operation{
		Name:       opDescribeEntitiesDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEntitiesDetectionJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEntitiesDetectionJobOutput{})
	return DescribeEntitiesDetectionJobRequest{Request: req, Input: input, Copy: c.DescribeEntitiesDetectionJobRequest}
}

// DescribeEntitiesDetectionJobRequest is the request type for the
// DescribeEntitiesDetectionJob API operation.
type DescribeEntitiesDetectionJobRequest struct {
	*aws.Request
	Input *types.DescribeEntitiesDetectionJobInput
	Copy  func(*types.DescribeEntitiesDetectionJobInput) DescribeEntitiesDetectionJobRequest
}

// Send marshals and sends the DescribeEntitiesDetectionJob API request.
func (r DescribeEntitiesDetectionJobRequest) Send(ctx context.Context) (*DescribeEntitiesDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEntitiesDetectionJobResponse{
		DescribeEntitiesDetectionJobOutput: r.Request.Data.(*types.DescribeEntitiesDetectionJobOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEntitiesDetectionJobResponse is the response type for the
// DescribeEntitiesDetectionJob API operation.
type DescribeEntitiesDetectionJobResponse struct {
	*types.DescribeEntitiesDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEntitiesDetectionJob request.
func (r *DescribeEntitiesDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
