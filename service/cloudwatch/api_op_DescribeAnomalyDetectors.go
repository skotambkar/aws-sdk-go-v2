// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const opDescribeAnomalyDetectors = "DescribeAnomalyDetectors"

// DescribeAnomalyDetectorsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Lists the anomaly detection models that you have created in your account.
// You can list all models in your account or filter the results to only the
// models that are related to a certain namespace, metric name, or metric dimension.
//
//    // Example sending a request using DescribeAnomalyDetectorsRequest.
//    req := client.DescribeAnomalyDetectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeAnomalyDetectors
func (c *Client) DescribeAnomalyDetectorsRequest(input *types.DescribeAnomalyDetectorsInput) DescribeAnomalyDetectorsRequest {
	op := &aws.Operation{
		Name:       opDescribeAnomalyDetectors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAnomalyDetectorsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAnomalyDetectorsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeAnomalyDetectorsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeAnomalyDetectorsRequest{Request: req, Input: input, Copy: c.DescribeAnomalyDetectorsRequest}
}

// DescribeAnomalyDetectorsRequest is the request type for the
// DescribeAnomalyDetectors API operation.
type DescribeAnomalyDetectorsRequest struct {
	*aws.Request
	Input *types.DescribeAnomalyDetectorsInput
	Copy  func(*types.DescribeAnomalyDetectorsInput) DescribeAnomalyDetectorsRequest
}

// Send marshals and sends the DescribeAnomalyDetectors API request.
func (r DescribeAnomalyDetectorsRequest) Send(ctx context.Context) (*DescribeAnomalyDetectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAnomalyDetectorsResponse{
		DescribeAnomalyDetectorsOutput: r.Request.Data.(*types.DescribeAnomalyDetectorsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAnomalyDetectorsResponse is the response type for the
// DescribeAnomalyDetectors API operation.
type DescribeAnomalyDetectorsResponse struct {
	*types.DescribeAnomalyDetectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAnomalyDetectors request.
func (r *DescribeAnomalyDetectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
