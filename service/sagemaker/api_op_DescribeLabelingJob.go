// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opDescribeLabelingJob = "DescribeLabelingJob"

// DescribeLabelingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets information about a labeling job.
//
//    // Example sending a request using DescribeLabelingJobRequest.
//    req := client.DescribeLabelingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeLabelingJob
func (c *Client) DescribeLabelingJobRequest(input *types.DescribeLabelingJobInput) DescribeLabelingJobRequest {
	op := &aws.Operation{
		Name:       opDescribeLabelingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeLabelingJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLabelingJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeLabelingJobMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLabelingJobRequest{Request: req, Input: input, Copy: c.DescribeLabelingJobRequest}
}

// DescribeLabelingJobRequest is the request type for the
// DescribeLabelingJob API operation.
type DescribeLabelingJobRequest struct {
	*aws.Request
	Input *types.DescribeLabelingJobInput
	Copy  func(*types.DescribeLabelingJobInput) DescribeLabelingJobRequest
}

// Send marshals and sends the DescribeLabelingJob API request.
func (r DescribeLabelingJobRequest) Send(ctx context.Context) (*DescribeLabelingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLabelingJobResponse{
		DescribeLabelingJobOutput: r.Request.Data.(*types.DescribeLabelingJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLabelingJobResponse is the response type for the
// DescribeLabelingJob API operation.
type DescribeLabelingJobResponse struct {
	*types.DescribeLabelingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLabelingJob request.
func (r *DescribeLabelingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
