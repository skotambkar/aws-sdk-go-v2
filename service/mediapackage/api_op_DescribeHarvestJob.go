// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
)

const opDescribeHarvestJob = "DescribeHarvestJob"

// DescribeHarvestJobRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Gets details about an existing HarvestJob.
//
//    // Example sending a request using DescribeHarvestJobRequest.
//    req := client.DescribeHarvestJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeHarvestJob
func (c *Client) DescribeHarvestJobRequest(input *types.DescribeHarvestJobInput) DescribeHarvestJobRequest {
	op := &aws.Operation{
		Name:       opDescribeHarvestJob,
		HTTPMethod: "GET",
		HTTPPath:   "/harvest_jobs/{id}",
	}

	if input == nil {
		input = &types.DescribeHarvestJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeHarvestJobOutput{})
	return DescribeHarvestJobRequest{Request: req, Input: input, Copy: c.DescribeHarvestJobRequest}
}

// DescribeHarvestJobRequest is the request type for the
// DescribeHarvestJob API operation.
type DescribeHarvestJobRequest struct {
	*aws.Request
	Input *types.DescribeHarvestJobInput
	Copy  func(*types.DescribeHarvestJobInput) DescribeHarvestJobRequest
}

// Send marshals and sends the DescribeHarvestJob API request.
func (r DescribeHarvestJobRequest) Send(ctx context.Context) (*DescribeHarvestJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHarvestJobResponse{
		DescribeHarvestJobOutput: r.Request.Data.(*types.DescribeHarvestJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHarvestJobResponse is the response type for the
// DescribeHarvestJob API operation.
type DescribeHarvestJobResponse struct {
	*types.DescribeHarvestJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHarvestJob request.
func (r *DescribeHarvestJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
