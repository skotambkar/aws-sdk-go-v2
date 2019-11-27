// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const opGetDistribution = "GetDistribution2019_03_26"

// GetDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the information about a distribution.
//
//    // Example sending a request using GetDistributionRequest.
//    req := client.GetDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetDistribution
func (c *Client) GetDistributionRequest(input *types.GetDistributionInput) GetDistributionRequest {
	op := &aws.Operation{
		Name:       opGetDistribution,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/distribution/{Id}",
	}

	if input == nil {
		input = &types.GetDistributionInput{}
	}

	req := c.newRequest(op, input, &types.GetDistributionOutput{})
	return GetDistributionRequest{Request: req, Input: input, Copy: c.GetDistributionRequest}
}

// GetDistributionRequest is the request type for the
// GetDistribution API operation.
type GetDistributionRequest struct {
	*aws.Request
	Input *types.GetDistributionInput
	Copy  func(*types.GetDistributionInput) GetDistributionRequest
}

// Send marshals and sends the GetDistribution API request.
func (r GetDistributionRequest) Send(ctx context.Context) (*GetDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDistributionResponse{
		GetDistributionOutput: r.Request.Data.(*types.GetDistributionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDistributionResponse is the response type for the
// GetDistribution API operation.
type GetDistributionResponse struct {
	*types.GetDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDistribution request.
func (r *GetDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
