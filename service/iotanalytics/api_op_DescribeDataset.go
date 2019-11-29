// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDescribeDataset = "DescribeDataset"

// DescribeDatasetRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a data set.
//
//    // Example sending a request using DescribeDatasetRequest.
//    req := client.DescribeDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeDataset
func (c *Client) DescribeDatasetRequest(input *types.DescribeDatasetInput) DescribeDatasetRequest {
	op := &aws.Operation{
		Name:       opDescribeDataset,
		HTTPMethod: "GET",
		HTTPPath:   "/datasets/{datasetName}",
	}

	if input == nil {
		input = &types.DescribeDatasetInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDatasetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeDatasetMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDatasetRequest{Request: req, Input: input, Copy: c.DescribeDatasetRequest}
}

// DescribeDatasetRequest is the request type for the
// DescribeDataset API operation.
type DescribeDatasetRequest struct {
	*aws.Request
	Input *types.DescribeDatasetInput
	Copy  func(*types.DescribeDatasetInput) DescribeDatasetRequest
}

// Send marshals and sends the DescribeDataset API request.
func (r DescribeDatasetRequest) Send(ctx context.Context) (*DescribeDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetResponse{
		DescribeDatasetOutput: r.Request.Data.(*types.DescribeDatasetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetResponse is the response type for the
// DescribeDataset API operation.
type DescribeDatasetResponse struct {
	*types.DescribeDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataset request.
func (r *DescribeDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
