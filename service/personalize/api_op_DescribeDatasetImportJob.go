// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opDescribeDatasetImportJob = "DescribeDatasetImportJob"

// DescribeDatasetImportJobRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes the dataset import job created by CreateDatasetImportJob, including
// the import job status.
//
//    // Example sending a request using DescribeDatasetImportJobRequest.
//    req := client.DescribeDatasetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeDatasetImportJob
func (c *Client) DescribeDatasetImportJobRequest(input *types.DescribeDatasetImportJobInput) DescribeDatasetImportJobRequest {
	op := &aws.Operation{
		Name:       opDescribeDatasetImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDatasetImportJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDatasetImportJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeDatasetImportJobMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDatasetImportJobRequest{Request: req, Input: input, Copy: c.DescribeDatasetImportJobRequest}
}

// DescribeDatasetImportJobRequest is the request type for the
// DescribeDatasetImportJob API operation.
type DescribeDatasetImportJobRequest struct {
	*aws.Request
	Input *types.DescribeDatasetImportJobInput
	Copy  func(*types.DescribeDatasetImportJobInput) DescribeDatasetImportJobRequest
}

// Send marshals and sends the DescribeDatasetImportJob API request.
func (r DescribeDatasetImportJobRequest) Send(ctx context.Context) (*DescribeDatasetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetImportJobResponse{
		DescribeDatasetImportJobOutput: r.Request.Data.(*types.DescribeDatasetImportJobOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetImportJobResponse is the response type for the
// DescribeDatasetImportJob API operation.
type DescribeDatasetImportJobResponse struct {
	*types.DescribeDatasetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDatasetImportJob request.
func (r *DescribeDatasetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
