// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const opDescribeMetricCollectionTypes = "DescribeMetricCollectionTypes"

// DescribeMetricCollectionTypesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the available CloudWatch metrics for Amazon EC2 Auto Scaling.
//
// The GroupStandbyInstances metric is not returned by default. You must explicitly
// request this metric when calling EnableMetricsCollection.
//
//    // Example sending a request using DescribeMetricCollectionTypesRequest.
//    req := client.DescribeMetricCollectionTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeMetricCollectionTypes
func (c *Client) DescribeMetricCollectionTypesRequest(input *types.DescribeMetricCollectionTypesInput) DescribeMetricCollectionTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeMetricCollectionTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeMetricCollectionTypesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeMetricCollectionTypesOutput{})
	return DescribeMetricCollectionTypesRequest{Request: req, Input: input, Copy: c.DescribeMetricCollectionTypesRequest}
}

// DescribeMetricCollectionTypesRequest is the request type for the
// DescribeMetricCollectionTypes API operation.
type DescribeMetricCollectionTypesRequest struct {
	*aws.Request
	Input *types.DescribeMetricCollectionTypesInput
	Copy  func(*types.DescribeMetricCollectionTypesInput) DescribeMetricCollectionTypesRequest
}

// Send marshals and sends the DescribeMetricCollectionTypes API request.
func (r DescribeMetricCollectionTypesRequest) Send(ctx context.Context) (*DescribeMetricCollectionTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMetricCollectionTypesResponse{
		DescribeMetricCollectionTypesOutput: r.Request.Data.(*types.DescribeMetricCollectionTypesOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMetricCollectionTypesResponse is the response type for the
// DescribeMetricCollectionTypes API operation.
type DescribeMetricCollectionTypesResponse struct {
	*types.DescribeMetricCollectionTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMetricCollectionTypes request.
func (r *DescribeMetricCollectionTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
