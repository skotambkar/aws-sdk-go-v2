// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetLoadBalancerMetricData = "GetLoadBalancerMetricData"

// GetLoadBalancerMetricDataRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about health metrics for your Lightsail load balancer.
//
//    // Example sending a request using GetLoadBalancerMetricDataRequest.
//    req := client.GetLoadBalancerMetricDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerMetricData
func (c *Client) GetLoadBalancerMetricDataRequest(input *types.GetLoadBalancerMetricDataInput) GetLoadBalancerMetricDataRequest {
	op := &aws.Operation{
		Name:       opGetLoadBalancerMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetLoadBalancerMetricDataInput{}
	}

	req := c.newRequest(op, input, &types.GetLoadBalancerMetricDataOutput{})
	return GetLoadBalancerMetricDataRequest{Request: req, Input: input, Copy: c.GetLoadBalancerMetricDataRequest}
}

// GetLoadBalancerMetricDataRequest is the request type for the
// GetLoadBalancerMetricData API operation.
type GetLoadBalancerMetricDataRequest struct {
	*aws.Request
	Input *types.GetLoadBalancerMetricDataInput
	Copy  func(*types.GetLoadBalancerMetricDataInput) GetLoadBalancerMetricDataRequest
}

// Send marshals and sends the GetLoadBalancerMetricData API request.
func (r GetLoadBalancerMetricDataRequest) Send(ctx context.Context) (*GetLoadBalancerMetricDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoadBalancerMetricDataResponse{
		GetLoadBalancerMetricDataOutput: r.Request.Data.(*types.GetLoadBalancerMetricDataOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoadBalancerMetricDataResponse is the response type for the
// GetLoadBalancerMetricData API operation.
type GetLoadBalancerMetricDataResponse struct {
	*types.GetLoadBalancerMetricDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoadBalancerMetricData request.
func (r *GetLoadBalancerMetricDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
