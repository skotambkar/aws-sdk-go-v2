// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/signer/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
)

const opDescribeSigningJob = "DescribeSigningJob"

// DescribeSigningJobRequest returns a request value for making API operation for
// AWS Signer.
//
// Returns information about a specific code signing job. You specify the job
// by using the jobId value that is returned by the StartSigningJob operation.
//
//    // Example sending a request using DescribeSigningJobRequest.
//    req := client.DescribeSigningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/DescribeSigningJob
func (c *Client) DescribeSigningJobRequest(input *types.DescribeSigningJobInput) DescribeSigningJobRequest {
	op := &aws.Operation{
		Name:       opDescribeSigningJob,
		HTTPMethod: "GET",
		HTTPPath:   "/signing-jobs/{jobId}",
	}

	if input == nil {
		input = &types.DescribeSigningJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSigningJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeSigningJobMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeSigningJobRequest{Request: req, Input: input, Copy: c.DescribeSigningJobRequest}
}

// DescribeSigningJobRequest is the request type for the
// DescribeSigningJob API operation.
type DescribeSigningJobRequest struct {
	*aws.Request
	Input *types.DescribeSigningJobInput
	Copy  func(*types.DescribeSigningJobInput) DescribeSigningJobRequest
}

// Send marshals and sends the DescribeSigningJob API request.
func (r DescribeSigningJobRequest) Send(ctx context.Context) (*DescribeSigningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSigningJobResponse{
		DescribeSigningJobOutput: r.Request.Data.(*types.DescribeSigningJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSigningJobResponse is the response type for the
// DescribeSigningJob API operation.
type DescribeSigningJobResponse struct {
	*types.DescribeSigningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSigningJob request.
func (r *DescribeSigningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
